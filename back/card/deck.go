package card

type Deck []string

func (deck Deck) Len() int {
	return len(deck)
}

// assumes that this deck has been sorted
func (deck Deck) GetLessOfCardType(cardType string) (string, bool) {
	for _, currentCard := range deck {
		if currentCard[:1] == cardType {
			return currentCard, true
		}
	}
	return "", false
}

func (deck Deck) HaseOfType(cardType string) bool {
	for _, currentCard := range deck {
		if currentCard[:1] == cardType {
			return true
		}
	}
	return false
}

func (deck Deck) Less(i, j int) bool {
	order1, order2 := intOrder(deck[i][0:1]), intOrder(deck[j][0:1])
	if order2 != order1 {
		return order2 > order1
	}
	compareRes := compareChars(deck[i][1:], deck[j][1:])
	return compareRes == -1

}

func (deck Deck) Swap(i, j int) {
	deck[i], deck[j] = deck[j], deck[i]
}

func intOrder(t string) int {
	switch t {
	case "h":
		return 0
	case "s":
		return 1
	case "d":
		return 2
	case "c":
		return 3
	}
	return -1
}
