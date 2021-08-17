package Handlers

import (
	"fmt"
	"github.com/farzad80rad/cards/back/Player"
	"github.com/farzad80rad/cards/back/card"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"sync"
)

type putCardInfo struct {
	Id uuid.UUID `json:"id" binding:"required"`
	CardToPut string `json:"CardToPut" binding:"required"`;
}


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func addNewWebSocket_B( w http.ResponseWriter , r * http.Request){
		fmt.Println("enter to websocket ")
		var playerInfo  Player.PLayerInfo
		upgrader.CheckOrigin = func(r *http.Request) bool { return true}
		conn , err := upgrader.Upgrade(w , r,nil)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		err = conn.ReadJSON(&playerInfo)
		if  err != nil{
			fmt.Println("errr" , err)
			return
		}
		fmt.Println("connnect to websocket ")


	thisBotGameInfo := botGameGroups[playerInfo.Id]
		for {
			select {
			case newPutInfo, valid := <- thisBotGameInfo.PutCardChan :
				if !valid {
					fmt.Println("closing socket and chan")
					close(thisBotGameInfo.PutCardChan)
					conn.Close()
					return;
				}
				if err := conn.WriteJSON(newPutInfo) ; err != nil {
					log.Println("errr:" , err)
				}
				fmt.Println("send message : " , newPutInfo );
			}
		}
}


func putCardHandler_B(c *gin.Context){
	var req putCardInfo;

	err := c.ShouldBindJSON(&req);
	if err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
	}
	fmt.Println("card to play " ,req.CardToPut);

	groupInfo := botGameGroups[req.Id]
	fmt.Println(groupInfo.CurrentPlayerIndex)
	if reflect.TypeOf(groupInfo.Group.Players[groupInfo.CurrentPlayerIndex]) == reflect.TypeOf(&Player.PLayerInfo{}){
		groupInfo.Group.NewCardToPlay = req.CardToPut
		groupInfo.Group.BotGroupCond.Signal()
		fmt.Println("singnal sent");
		c.JSON(http.StatusOK , gin.H{"status": true } )
	}else{
		groupInfo.Group.NewCardToPlay = ""
		c.JSON(http.StatusOK , gin.H{"status": false } )
	}

}

func startGame_B(gameInfo *BotGameInfo){
	for gameInfo.CurrentPlayerIndex=0 ; ;gameInfo.CurrentPlayerIndex++ {
		gameInfo.CurrentPlayerIndex = gameInfo.CurrentPlayerIndex % 4;
		currentPlayer := gameInfo.Group.Players[gameInfo.CurrentPlayerIndex];
		fmt.Println(reflect.TypeOf(currentPlayer));
		fmt.Println(reflect.TypeOf(Player.PLayerInfo{}));
		if reflect.TypeOf(currentPlayer) == reflect.TypeOf(&Player.BotInfo{}) {
		 	cardToPlay := currentPlayer.PlayCard(gameInfo.OnBoardCards)
			gameInfo.OnBoardCards = append(gameInfo.OnBoardCards, cardToPlay)
			gameInfo.PutCardChan <- putCardInfo{CardToPut: cardToPlay, Id: currentPlayer.(*Player.BotInfo).Id}

		}else if reflect.TypeOf(currentPlayer) == reflect.TypeOf(&Player.PLayerInfo{}){
			gameInfo.Group.BotGroupCond.L.Lock()
			for gameInfo.Group.NewCardToPlay == "" {
				gameInfo.Group.BotGroupCond.Wait()
				fmt.Println("awaked")
			}
			gameInfo.Group.BotGroupCond.L.Unlock()
			gameInfo.OnBoardCards = append(gameInfo.OnBoardCards, gameInfo.Group.NewCardToPlay)
			gameInfo.Group.NewCardToPlay = ""
		}else{
			fmt.Println("something wrong!")
		}
		fmt.Println("index: " , gameInfo.CurrentPlayerIndex , "  onbouard" , gameInfo.OnBoardCards)

		if len(gameInfo.OnBoardCards) >= 4 {
			var max = "s2";
			var maxIndex = 0;
			for index , value := range gameInfo.OnBoardCards {
				compaerRes := card.Compare(max,value,gameInfo.OnBoardCards[0][:1],"h")
				fmt.Println(compaerRes)
				if (compaerRes < 0){
					maxIndex = index;
					max = value;
				}
			}
			gameInfo.CurrentPlayerIndex = (gameInfo.CurrentPlayerIndex + 1) % 4;
			gameInfo.CurrentPlayerIndex += maxIndex-1 ;

			gameInfo.OnBoardCards = make([]string,0)
		}
	}
}



func initGame_B(c *gin.Context){
	var newPlayer Player.PLayerInfo;
	err := c.ShouldBindJSON(&newPlayer)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return;
	}

	newPlayer.Id = uuid.New()
	newTotaldeck := card.GetCardsSuffuledArray()
	newPlayer.Deck = newTotaldeck[:13]
	bots := makeNewBotGroup(newTotaldeck[13:])


	var newBotGroup  BotGroup;
	newBotGroup.BotGroupCond = sync.NewCond(&sync.Mutex{})
	newBotGroup.Players = []Player.Player{ &newPlayer , &bots[0] ,&bots[1] ,&bots[2]  } ;


	var botGroupToAdd BotGameInfo;
	botGroupToAdd.CurrentPlayerIndex = 0;
	botGroupToAdd.OnBoardCards = make([]string,0)
	botGroupToAdd.Group = &newBotGroup
	botGroupToAdd.PutCardChan = make(chan putCardInfo,1);

	botGameGroups[newPlayer.Id] = &botGroupToAdd
	c.JSON(http.StatusOK, gin.H{
		"player": newPlayer,
		"bots": bots,
	})
	go startGame_B(&botGroupToAdd)
	return
}

func makeNewBotGroup( totalDeck []string ) []Player.BotInfo{
	new3Bots := make([]Player.BotInfo,3)
	for i:=0 ; i<3 ; i++ {
		new3Bots[i].Name = "Bot " + strconv.Itoa(i)
		new3Bots[i].Id = uuid.New()
		new3Bots[i].Deck = totalDeck[:13]
		totalDeck = totalDeck[13:]
	}
	return new3Bots
}