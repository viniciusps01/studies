package deck

import (
	"math/rand"
	"strings"
)

type card struct {
	Value int    `json:"value"`
	Suit  string `json:"suit"`
}

func NewCard(value int, suit string) card {
	card := card{
		Value: value,
		Suit:  suit,
	}

	return card
}

func (c *card) ToString() string {
	valueName := valueName(c.Value)
	data := []string{valueName, " of ", c.Suit}
	return strings.Join(data, "")
}

func shuffleCards(cards *[]card) {
	c := *cards
	rand.Shuffle(len(*cards), func(i int, j int) {
		c[i], c[j] = c[j], c[i]
	})
}

func valueName(value int) string {
	name := "invalid number"

	switch value {
	case 1:
		name = "Ace"
	case 2:
		name = "Two"
	case 3:
		name = "Three"
	case 4:
		name = "Four"
	case 5:
		name = "Five"
	case 6:
		name = "Six"
	case 7:
		name = "Seven"
	case 8:
		name = "Eight"
	case 9:
		name = "Nine"
	case 10:
		name = "Ten"
	case 11:
		name = "Jack"
	case 12:
		name = "Queen"
	case 13:
		name = "King"
	}

	return name
}
