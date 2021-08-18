package Player

import (
	"github.com/farzad80rad/cards/back/card"
	"github.com/google/uuid"
)

type BotInfo struct {
	Name string `json:"username" binding:"-"`;
	Id uuid.UUID `json:"id" binding:"required"`
	Deck card.Deck `json:"deck" binding:"-"`;
}