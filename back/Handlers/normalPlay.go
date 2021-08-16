package Handlers

import (
	"github.com/farzad80rad/cards/back/Player"
	"github.com/farzad80rad/cards/back/card"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

var muPlayersSlice sync.Mutex;

func handelNewPlayer( c *gin.Context) {
	var newPlayer Player.PLayerInfo;
	err := c.ShouldBindJSON(&newPlayer)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return;
	}
	addNewPlayer(newPlayer.UserName)
}

func addNewPlayer(name string){
	pendingPlayers = append(pendingPlayers,name );
	muPlayersSlice.Lock()
	if len(pendingPlayers) >= 4 {
		newPlyersToAdd := make([]Player.PLayerInfo,4);
		for i :=0 ;i<4; i++{
			newPlyersToAdd[i].Id = uuid.New()
			newPlyersToAdd[i].Deck = []string{card.C2,card.Ca,card.D4,card.D8,card.Dj,card.H9,card.S2}
			newPlyersToAdd[i].UserName = pendingPlayers[i]
		}
		newGroup := Group{Players : newPlyersToAdd }
		pendingPlayers = pendingPlayers[4:];
		groups = append(groups, newGroup)
	}
	muPlayersSlice.Unlock()
}
