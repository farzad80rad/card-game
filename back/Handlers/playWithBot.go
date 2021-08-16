package Handlers

import (
	"fmt"
	"github.com/farzad80rad/cards/back/Player"
	"github.com/farzad80rad/cards/back/card"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"reflect"
	"sync"
)

type putCard struct {
	Id uuid.UUID `json:"id" binding:"required"`
	CardToPut string `json:"CardToPut" binding:"required"`;
}

func putCardHandler_B(c *gin.Context){
	var req putCard;

	err := c.ShouldBindJSON(&req);
	if err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
	}
	fmt.Println("card to play " ,req.CardToPut);

	groupInfo := botGameGroups[req.Id]
	fmt.Println("singnal sent");
	fmt.Println(groupInfo.CurrentPlayerIndex)
	if reflect.TypeOf(groupInfo.Group.Players[groupInfo.CurrentPlayerIndex]) == reflect.TypeOf(&Player.PLayerInfo{}){
		groupInfo.Group.NewCardToPlay = req.CardToPut
		groupInfo.Group.BotGroupCond.Signal()
		c.JSON(http.StatusOK , gin.H{"status": true } )
	}else{
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

		}else if reflect.TypeOf(currentPlayer) == reflect.TypeOf(&Player.PLayerInfo{}){
			gameInfo.Group.BotGroupCond.L.Lock()
			for gameInfo.Group.NewCardToPlay == "" {
				gameInfo.Group.BotGroupCond.Wait()
				fmt.Println("awaked")
				fmt.Println(gameInfo.Group.NewCardToPlay);
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
		new3Bots[i].Name = "Bot1"
		new3Bots[i].Deck = totalDeck[:13]
		totalDeck = totalDeck[13:]
	}
	return new3Bots
}