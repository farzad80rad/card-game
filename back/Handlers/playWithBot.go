package Handlers

import (
	"fmt"
	"github.com/farzad80rad/cards/back/Player"
	"github.com/farzad80rad/cards/back/card"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func handelPlayWithBot(c *gin.Context){
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
	newBotGroup.Player = newPlayer;
	newBotGroup.Bots = bots;
	botGameGroups[newPlayer.Id] = newBotGroup
	c.JSON(http.StatusOK, gin.H{
		"player": newPlayer,
		"bots": bots,
	})
	fmt.Println(newPlayer)
	fmt.Println(bots)
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