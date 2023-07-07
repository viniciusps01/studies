package deck

import "testing"

func TestNew(t *testing.T) {
	d1 := New()
	d2 := New()

	length := len(d1.Cards)

	if length != 52 {
		t.Errorf("Expected deck length of 52, but got %v", length)
	}

	sameOrder := true

	for i1, v1 := range d1.Cards {
		v2 := d2.Cards[i1]

		if v1 != v2 {
			sameOrder = false
			break
		}
	}

	if sameOrder {
		t.Errorf("Expected new deck to have shuffled cards but didn't got that")
	}
}

func TestShuffle(t *testing.T) {
	d1 := New()
	d2 := New()

	sameOrder := true

	for i1, v1 := range d1.Cards {
		v2 := d2.Cards[i1]

		if v1 != v2 {
			sameOrder = false
			break
		}
	}

	if sameOrder {
		t.Errorf("Expected new deck to have shuffled cards but didn't got that")
	}
}

func TestAdd(t *testing.T) {
	deck := New()

	previousLength := len(deck.Cards)
	card := NewCard(1, "suit")
	deck.Add(card)

	expectedLength := previousLength + 1
	currentLength := len(deck.Cards)

	if currentLength != expectedLength {
		t.Errorf("Expected length to be %v, but got %b", expectedLength, currentLength)
	}
}

func TestDeal(t *testing.T) {
	deck := New()
	previousDeckLength := len(deck.Cards)

	quantity := 5
	hand := deck.Deal(quantity)
	currentDeckLength := len(deck.Cards)

	if len(hand.Cards) != quantity {
		t.Errorf("Expected hand to have %v cards, but got %v ones", quantity, len(hand.Cards))
	}

	expectedDeckLength := previousDeckLength - 5
	if currentDeckLength != expectedDeckLength {
		t.Errorf("Expected deck to have %v cards, but got %v ones", expectedDeckLength, currentDeckLength)
	}
}
