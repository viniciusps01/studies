package deck

import "fmt"

type hand struct {
	Cards []card `json:"cards"`
}

func (h *hand) Shuffle() {
	shuffleCards(&h.Cards)
}

func (h hand) Print() {
	for _, card := range h.Cards {
		fmt.Println(card.ToString())
	}
}
