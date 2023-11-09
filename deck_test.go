package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 54 {
		t.Errorf("Expected deck length of 54, but got %v.", len(deck))
	}
	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v.", deck[0])
	}

	if deck[len(deck)-3] != "Ten of Clubs" {
		t.Errorf("Expected last card of Ten of Clubs, but got %v.", deck[len(deck)-1])
	}

	if deck[len(deck)-1] != "Black and White Joker" {
		t.Errorf("Expected last card of Black and White Joker, but got %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 54 {
		t.Errorf("Expected 54 cards in deck, but got %v.", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
