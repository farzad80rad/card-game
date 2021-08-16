package Handlers

import (
	"github.com/farzad80rad/cards/back/Player"
	"github.com/google/uuid"
	"sync"
)

type Group struct {
	 Players []Player.PLayerInfo;
}

type BotGroup struct {
	Players []Player.Player ;
	BotGroupCond *sync.Cond ;
	NewCardToPlay string;
}

type BotGameInfo struct {
	Group *BotGroup `json:"botsAndPlayer" binding:"-"`
	CurrentPlayerIndex int  `json:"index" binding:"-"` // 0 for selfPlayer and others for bots
	OnBoardCards []string  `json:"deck" binding:"-"`
}

var pendingPlayers []string ;
var groups []Group;
var botGameGroups map[uuid.UUID]*BotGameInfo

func initSys(){
	pendingPlayers = make([]string, 0);
	botGameGroups = make(map[uuid.UUID]*BotGameInfo, 0);
	groups = make([]Group,0);
}


