package card

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	Ca  = "ca"
	C2  = "c2"
	C3  = "c3"
	C4  = "c4"
	C5  = "c5"
	C6  = "c6"
	C7  = "c7"
	C8  = "c8"
	C9  = "c9"
	C10 = "c10"
	Cj  = "cj"
	Cq  = "cq"
	Ck  = "ck"

	Da  = "da"
	D2  = "d2"
	D3  = "d3"
	D4  = "d4"
	D5  = "d5"
	D6  = "d6"
	D7  = "d7"
	D8  = "d8"
	D9  = "d9"
	D10 = "d10"
	Dj  = "dj"
	Dq  = "dq"
	Dk  = "dk"

	Ha  = "ha"
	H2  = "h2"
	H3  = "h3"
	H4  = "h4"
	H5  = "h5"
	H6  = "h6"
	H7  = "h7"
	H8  = "h8"
	H9  = "h9"
	H10 = "h10"
	Hj  = "hj"
	Hq  = "hq"
	Hk  = "hk"

	Sa  = "sa"
	S2  = "s2"
	S3  = "s3"
	S4  = "s4"
	S5  = "s5"
	S6  = "s6"
	S7  = "s7"
	S8  = "s8"
	S9  = "s9"
	S10 = "s10"
	Sj  = "sj"
	Sq  = "sq"
	Sk  = "sk"
)

func GetCardsSuffuledArray() []string {
	newDeck := []string{
		Ca, C2, C3, C4, C5, C6, C7, C8, C9, C10, Cj, Cq, Ck,
		Ha, H2, H3, H4, H5, H6, H7, H8, H9, H10, Hj, Hq, Hk,
		Sa, S2, S3, S4, S5, S6, S7, S8, S9, S10, Sj, Sq, Sk,
		Da, D2, D3, D4, D5, D6, D7, D8, D9, D10, Dj, Dq, Dk,
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newDeck), func(i, j int) {
		newDeck[i], newDeck[j] = newDeck[j], newDeck[i]
	})
	return newDeck
}

func Compare(card1, card2, playingCard, hokm string) int {
	fmt.Println("compare ", card1, " and ", card2, " playing card is ", playingCard, "hokm is ", hokm)
	if strings.HasPrefix(card1, playingCard) && strings.HasPrefix(card2, playingCard) {
		return compareChars(card1[1:], card2[1:])
	}
	if strings.HasPrefix(card1, hokm) && strings.HasPrefix(card2, hokm) {
		return compareChars(card1[1:], card2[1:])
	}
	if strings.HasPrefix(card2, hokm) {
		return -1
	}
	if strings.HasPrefix(card1, hokm) {
		return 1
	}
	if strings.HasPrefix(card2, playingCard) {
		return -1
	}
	if strings.HasPrefix(card1, playingCard) {
		return 1
	}
	return 0
}

func compareChars(type1, type2 string) int {
	switch {
	case type1 == type2:
		return 0
	case type1 == "a":
		return 1
	case type2 == "a":
		return -1
	case type1 == "k":
		return 1
	case type2 == "k":
		return -1
	case type1 == "q":
		return 1
	case type2 == "q":
		return -1
	case type1 == "j":
		return 1
	case type2 == "j":
		return -1
	case type1 == "10":
		return 1
	case type2 == "10":
		return -1
	default:
		return strings.Compare(type1, type2)
	}

}
