package Handlers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/farzad80rad/cards/back/Player"
	"github.com/farzad80rad/cards/back/card"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type putCardInfo struct {
	Id        uuid.UUID `json:"id" binding:"required"`
	CardToPut string    `json:"CardToPut" binding:"required"`
}

type websocketSendingInfo struct {
	Id          uuid.UUID `json:"id" `
	CardToPut   string    `json:"CardToPut" `
	Type        string    `json:"type"`
	SelfTeamWin bool      `json:"SelfTeamWin"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func addNewWebSocket_B(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter to websocket ")
	var playerInfo Player.PLayerInfo
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	err = conn.ReadJSON(&playerInfo)
	if err != nil {
		fmt.Println("errr", err)
		return
	}
	fmt.Println("connnect to websocket ")
	go serveSocket_B(playerInfo.Id, conn)
}

func serveSocket_B(id uuid.UUID, conn *websocket.Conn) {

	thisBotGameInfo := botGameGroups[id]
	for {
		select {
		case newPutInfo := <-thisBotGameInfo.PutCardChan:
			if err := conn.WriteJSON(websocketSendingInfo{Id: newPutInfo.Id, CardToPut: newPutInfo.CardToPut, Type: "put card"}); err != nil {
				log.Println("errr:", err)
			}
			fmt.Println("send message : ", newPutInfo)

		case selfWin := <-thisBotGameInfo.CleanTableChan:
			if selfWin {
				if err := conn.WriteJSON(websocketSendingInfo{Type: "clean table", SelfTeamWin: true}); err != nil {
					log.Println("errr:", err)
				}
			} else {
				if err := conn.WriteJSON(websocketSendingInfo{Type: "clean table", SelfTeamWin: false}); err != nil {
					log.Println("errr:", err)
				}
			}

			fmt.Println("should clean table")
		}
	}
}

func putCardHandler(c *gin.Context) {
	var req putCardInfo

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	groupInfo := botGameGroups[req.Id]
	fmt.Println(groupInfo.CurrentPlayerIndex)
	currentPlayer := groupInfo.Group.Players[groupInfo.CurrentPlayerIndex]
	if reflect.TypeOf(currentPlayer) == reflect.TypeOf(&Player.PLayerInfo{}) && (currentPlayer.(*Player.PLayerInfo).Id == req.Id) {
		groupInfo.Group.NewCardToPlay = req.CardToPut
		groupInfo.Group.BotGroupCond.Signal()
		c.JSON(http.StatusOK, gin.H{"status": true})
	} else {
		groupInfo.Group.NewCardToPlay = ""
		c.JSON(http.StatusOK, gin.H{"status": false})
	}

}

func startGame(gameInfo *BotGameInfo) {
	team1Wins := 0
	team2Wins := 0
	hokm := "h"
	for gameInfo.CurrentPlayerIndex = 0; team1Wins < 7 && team2Wins < 7; gameInfo.CurrentPlayerIndex++ {
		gameInfo.CurrentPlayerIndex = gameInfo.CurrentPlayerIndex % 4
		currentPlayer := gameInfo.Group.Players[gameInfo.CurrentPlayerIndex]

		if reflect.TypeOf(currentPlayer) == reflect.TypeOf(&Player.BotInfo{}) {
			// bot turn
			cardToPlay := currentPlayer.PlayCard(gameInfo.OnBoardCards, hokm)
			gameInfo.OnBoardCards = append(gameInfo.OnBoardCards, cardToPlay)
			gameInfo.OnBoardCardsPutters[cardToPlay] = currentPlayer.(*Player.BotInfo).Id
			gameInfo.PutCardChan <- putCardInfo{CardToPut: cardToPlay, Id: currentPlayer.(*Player.BotInfo).Id}

		} else if reflect.TypeOf(currentPlayer) == reflect.TypeOf(&Player.PLayerInfo{}) {
			// player turn
			gameInfo.Group.BotGroupCond.L.Lock()
			for gameInfo.Group.NewCardToPlay == "" {
				gameInfo.Group.BotGroupCond.Wait()
			}
			gameInfo.Group.BotGroupCond.L.Unlock()
			gameInfo.OnBoardCards = append(gameInfo.OnBoardCards, gameInfo.Group.NewCardToPlay)
			gameInfo.OnBoardCardsPutters[gameInfo.Group.NewCardToPlay] = currentPlayer.(*Player.PLayerInfo).Id
			gameInfo.Group.NewCardToPlay = ""
		} else {
			fmt.Println("something wrong!")
		}

		if len(gameInfo.OnBoardCards) >= 4 {
			// end of each turn
			var max = "s2" // just as default. must be somthing that could be the less valuble card in the deck.
			var maxIndex = 0
			for index, value := range gameInfo.OnBoardCards {
				compaerRes := card.Compare(max, value, gameInfo.OnBoardCards[0][:1], hokm)
				fmt.Println(compaerRes)
				if compaerRes < 0 {
					maxIndex = index
					max = value
				}
			}
			fmt.Println("max : ", max)
			gameInfo.CurrentPlayerIndex = (gameInfo.CurrentPlayerIndex + 1) % 4
			gameInfo.CurrentPlayerIndex += maxIndex - 1
			gameInfo.OnBoardCards = make(card.Deck, 0)
			time.Sleep(1500 * time.Millisecond)
			players := gameInfo.Group.Players
			if gameInfo.OnBoardCardsPutters[max] == players[0].GetId() || gameInfo.OnBoardCardsPutters[max] == players[2].GetId() {
				fmt.Println("self win")
				gameInfo.CleanTableChan <- true
				team1Wins += 1

			} else {
				team2Wins += 1
				gameInfo.CleanTableChan <- false
			}

		}
	}
}

func initGame_B(c *gin.Context) {
	var newPlayer Player.PLayerInfo
	err := c.ShouldBindJSON(&newPlayer)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newPlayer.Id = uuid.New()
	newTotaldeck := card.GetCardsSuffuledArray()
	newPlayer.Deck = newTotaldeck[:13]
	sort.Sort(newPlayer.Deck)
	bots := makeNewBotGroup(newTotaldeck[13:])

	var newBotGroup BotGroup
	newBotGroup.BotGroupCond = sync.NewCond(&sync.Mutex{})
	newBotGroup.Players = []Player.Player{&newPlayer, &bots[0], &bots[1], &bots[2]}

	var botGroupToAdd BotGameInfo
	botGroupToAdd.CurrentPlayerIndex = 0
	botGroupToAdd.OnBoardCards = make([]string, 0)
	botGroupToAdd.Group = &newBotGroup
	botGroupToAdd.OnBoardCardsPutters = make(map[string]uuid.UUID, 0)
	botGroupToAdd.PutCardChan = make(chan putCardInfo, 1)
	botGroupToAdd.CleanTableChan = make(chan bool, 1)

	botGameGroups[newPlayer.Id] = &botGroupToAdd
	c.JSON(http.StatusOK, gin.H{
		"player": newPlayer,
		"bots":   bots,
	})
	go startGame(&botGroupToAdd)
}

func makeNewBotGroup(totalDeck card.Deck) []Player.BotInfo {
	new3Bots := make([]Player.BotInfo, 3)
	for i := 0; i < 3; i++ {
		new3Bots[i].Name = "Bot " + strconv.Itoa(i)
		new3Bots[i].Id = uuid.New()
		new3Bots[i].Deck = totalDeck[:13]
		sort.Sort(new3Bots[i].Deck)
		totalDeck = totalDeck[13:]
	}
	return new3Bots
}
