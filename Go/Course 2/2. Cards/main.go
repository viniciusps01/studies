package main

import (
	deck "app/modules/deck"
	"fmt"
	"os"
)

func main() {
	var option string
	var cardsDeck deck.Deck
	const fileName = "deck.txt"

	for {
		showMenu()
		option = readOption()

		switch option {
		case "1":
			cardsDeck = deck.New()
		case "2":
			cardsDeck.Print()
		case "3":
			cardsDeck.Shuffle()
		case "4":
			cardsDeck.Deal(5)
		case "5":
			err := saveDeck(cardsDeck, fileName)

			showError(err)

		case "6":
			res, err := readDeck(fileName)

			if err != nil {
				showError(err)
				continue
			}

			cardsDeck = *res
		}
	}
}

func showError(err error) {
	if err == nil {
		return
	}

	fmt.Println("Error:", err)
}

func readOption() string {
	var option string
	fmt.Scan(&option)
	return option
}

func showMenu() {
	fmt.Println("===============================")
	fmt.Println("1) Create new deck")
	fmt.Println("2) Print deck")
	fmt.Println("3) Shuffle deck")
	fmt.Println("4) Heal")
	fmt.Println("5) Save deck")
	fmt.Println("6) Read saved deck")
	fmt.Println("===============================")
	fmt.Println()
}

func readDeck(fileName string) (*deck.Deck, error) {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	deck := deck.NewDeckFromBytes(bytes)

	return &deck, nil
}

func saveDeck(deck deck.Deck, fileName string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	_, writeErr := file.Write(deck.ToBytes())

	file.Close()

	if writeErr != nil {
		return writeErr
	}

	return nil
}
