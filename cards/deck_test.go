package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 4 {
		t.Errorf("Expected deck lenght of 4, but got %v", len(d))
	}

	if d[0] != "One of Apple" {
		t.Errorf("Expected first card is One of Apple, but got %v", d[0])
	}
}
