package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of length 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// remove any "dirty" files
	os.Remove("_decktesting")

	//create a new deck
	deck := newDeck()
	// save deck to hard drive
	deck.saveToFile("_decktesting")
	// load the file from the disk
	loadedDeck := newDeckFromFile("_deckTesting")
	// make sure that the deck testing file loaded
	if len(loadedDeck) != 16 {
		t.Errorf("Expecting 16 cards in deck got %v", len(loadedDeck))
	}
	// do some cleanup for our test [delete the test file we created]
	os.Remove("_decktesting")
}
