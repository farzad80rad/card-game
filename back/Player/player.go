package Player

import "github.com/google/uuid"

type Player interface {
	PlayCard([]string) string
}

type PLayerInfo struct  {
	UserName string `json:"username" binding:"-"`;
	Id uuid.UUID `json:"id" binding:"-"`
	Deck []string `json:"deck" binding:"-"`;
}

