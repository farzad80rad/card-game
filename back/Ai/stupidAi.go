package Ai

import "math/rand"

type AiInfo struct  {
	name string;
	deck []string;
}


func (ai *AiInfo) PlayCard(perviosCards []string) string{
	randomIndex := rand.Intn(len(ai.deck))
	pick := ai.deck[randomIndex]
	ai.deck =  append(ai.deck[:randomIndex], ai.deck[randomIndex+1:]...)
	return pick
}
