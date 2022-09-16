package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected lenght: 40, but got %d", len(d))
	}
}
