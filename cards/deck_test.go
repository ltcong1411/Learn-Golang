package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 4 {
		t.Errorf("Expected deck lenght of 4, but got %v", len(d))
	}

	if d[0] != "One of Apple" {
		t.Errorf("Expected first card is One of Apple, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 4 {
		t.Errorf("Expected 4 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
