package Player

import "math/rand"



func (ai *BotInfo) PlayCard(perviosCards []string) string{
	randomIndex := rand.Intn(len(ai.Deck))
	pick := ai.Deck[randomIndex]
	ai.Deck =  append(ai.Deck[:randomIndex], ai.Deck[randomIndex+1:]...)
	return pick
}

