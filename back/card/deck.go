package card

type Deck []string

func (deck Deck) Len() int {
	return len(deck)
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
