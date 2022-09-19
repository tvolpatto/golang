package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected lenght: 40, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element: Ace of Spades, but got %s", d[0])

	}

	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("Expected last element: Ace of Spades, but got %s", d[len(d)-1])

	}
}

func TestSaveToFileAndOpenDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 40 {
		t.Errorf("Expected lenght: 40, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
