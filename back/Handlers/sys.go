package Handlers

import (
	"sync"

	"github.com/farzad80rad/cards/back/Player"
	"github.com/farzad80rad/cards/back/card"
	"github.com/google/uuid"
)

type Group struct {
	Players []Player.PLayerInfo
}

type BotGroup struct {
	Players       []Player.Player
	BotGroupCond  *sync.Cond
	NewCardToPlay string
}

type MessageInfo struct {
	Sender   string    `json:"sender" binding:"required"`
	SenderID uuid.UUID `json:"senderId" binding:"required"`
	Message  string    `json:"message" binding:"required"`
}

type BotGameInfo struct {
	Group               *BotGroup            `json:"botsAndPlayer" binding:"-"`
	CurrentPlayerIndex  int                  `json:"index" binding:"-"` // 0 for selfPlayer and others for bots
	OnBoardCards        card.Deck            `json:"deck" binding:"-"`
	OnBoardCardsPutters map[string]uuid.UUID // this is for take track of which card was played by who.
	PutCardChan         chan putCardInfo
	CleanTableChan      chan bool
	MessageChan         chan MessageInfo
}

var pendingPlayers []string
var groups []Group
var botGameGroups map[uuid.UUID]*BotGameInfo

func initSys() {
	pendingPlayers = make([]string, 0)
	botGameGroups = make(map[uuid.UUID]*BotGameInfo, 0)
	groups = make([]Group, 0)
}
