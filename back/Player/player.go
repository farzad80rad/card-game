package Player

import (
	"github.com/farzad80rad/cards/back/card"
	"github.com/google/uuid"
)

type Player interface {
	PlayCard([]string, string) string
	GetId() uuid.UUID
}

type PLayerInfo struct {
	UserName string    `json:"username" binding:"-"`
	Id       uuid.UUID `json:"id" binding:"-"`
	Deck     card.Deck `json:"deck" binding:"-"`
}

func (player *PLayerInfo) PlayCard(perviosCards []string, hokm string) string {
	return ""
}

func (player *PLayerInfo) GetId() uuid.UUID {
	return player.Id
}
