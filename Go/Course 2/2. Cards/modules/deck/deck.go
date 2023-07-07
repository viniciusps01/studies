package deck

import (
	"encoding/json"
	"fmt"
)

type Deck struct {
	Cards []card `json:"cards"`
}

func New() Deck {
	deck := Deck{}
	deck.populateDeck()
	deck.Shuffle()
	return deck
}

func (d *Deck) populateDeck() {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, name := range cardSuits {
		for i := 1; i <= 13; i++ {
			card := NewCard(i, name)
			d.Add(card)
		}

	}
}

func NewDeckFromBytes(bytes []byte) Deck {
	var deck Deck

	json.Unmarshal(bytes, &deck)

	return deck
}

func (d Deck) ToBytes() []byte {
	bytes, _ := json.Marshal(d)
	return bytes
}

func (deck *Deck) Shuffle() {
	shuffleCards(&deck.Cards)
}

func (deck *Deck) Add(card card) {
	deck.Cards = append(deck.Cards, card)
}

func (deck *Deck) Deal(quantity int) hand {
	hand := hand{
		Cards: deck.Cards[:quantity],
	}

	deck.Cards = deck.Cards[quantity:]

	return hand
}

func (deck *Deck) Print() {
	for _, card := range deck.Cards {
		fmt.Println(card.ToString())
	}
}
