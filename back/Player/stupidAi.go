package Player

import (
	"math/rand"
	"time"
)



func (ai *BotInfo) PlayCard(perviosCards []string) string{
	time.Sleep(1150 * time.Millisecond)
	randomIndex := rand.Intn(len(ai.Deck))
	pick := ai.Deck[randomIndex]
	ai.Deck =  append(ai.Deck[:randomIndex], ai.Deck[randomIndex+1:]...)
	return pick
}

