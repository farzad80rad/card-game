package Player

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/farzad80rad/cards/back/card"
	"github.com/google/uuid"
)

func (ai *BotInfo) GetId() uuid.UUID {
	return ai.Id
}
func (ai *BotInfo) GetName() string {
	return ai.Name
}

func (ai *BotInfo) PlayCard(perviosCards []string, hokm string) string {
	time.Sleep(1150 * time.Millisecond)
	randomIndex := rand.Intn(len(ai.Deck))
	var pick string
	fmt.Println("cards of bot: ", ai.Deck, "   pervios cards : ", perviosCards)
	switch len(perviosCards) {
	case 0:
		pick = ai.Deck[randomIndex]
	default:
		for _, currentCard := range ai.Deck {

			if card.Compare(currentCard, perviosCards[len(perviosCards)-1], perviosCards[0][:1], hokm) > 0 {
				if !(currentCard[:1] == hokm && ai.Deck.HaseOfType(perviosCards[0][:1])) {
					pick = currentCard
				}
			}
		}
		if pick == "" {
			if smallestCard, found := ai.Deck.GetLessOfCardType(perviosCards[0][:1]); found {
				pick = smallestCard
			} else {
				pick = ai.Deck[randomIndex]

			}
		}

	}
	for index, cardName := range ai.Deck {
		if cardName == pick {
			ai.Deck = append(ai.Deck[:index], ai.Deck[index+1:]...)
			break
		}
	}
	return pick
}
