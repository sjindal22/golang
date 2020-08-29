package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Size of deck expected to be 16 but got %v ", len(d))
	}

	if d[0] != "Ace of Club" {
		t.Errorf("Expected Ace of Club but got %v ", d[0])
	}

	if d[len(d)-1] != "Three of Spades" {
		t.Errorf("Expected Three of Spades but got %v ", d[len(d)-1])
	}
}

// Revise lecture 29 and 24 again. The function are missing.
