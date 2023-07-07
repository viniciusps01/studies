package main

import (
	deck "app/modules/deck"
	"os"
	"testing"
)

func TestSaveDeckToFileAndLoadDeckFromFile(t *testing.T) {
	fileName := "_deck_testing.temp"

	os.Remove(fileName)

	newDeck := deck.New()

	saveErr := saveDeck(newDeck, fileName)

	if saveErr != nil {
		t.Errorf("Failed to save deck")
	}

	loadedDeck, loadErr := readDeck(fileName)

	if loadErr != nil {
		t.Errorf("Failed to load deck")
	}

	if string(newDeck.ToBytes()) != string(loadedDeck.ToBytes()) {
		t.Errorf("Expected saved deck loaded one to match but failed")
	}

	os.Remove(fileName)
}
