package Player

import "github.com/google/uuid"

type BotInfo struct {
	Name string `json:"username" binding:"-"`;
	Id uuid.UUID `json:"id" binding:"required"`
	Deck []string `json:"deck" binding:"-"`;
}