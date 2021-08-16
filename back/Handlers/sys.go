package Handlers

import (
	"github.com/farzad80rad/cards/back/Player"
	"github.com/google/uuid"
)

type Group struct {
	 Players []Player.PLayerInfo;
}

type BotGroup struct {
	Player Player.PLayerInfo;
	Bots []Player.BotInfo;
}

var pendingPlayers []string ;
var groups []Group;
var botGameGroups map[uuid.UUID]BotGroup

func initSys(){
	pendingPlayers = make([]string, 0);
	botGameGroups = make(map[uuid.UUID]BotGroup, 0);
	groups = make([]Group,0);
}


