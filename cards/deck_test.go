package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(cards))
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	fileName := "deckTestFile"
	os.Remove(fileName)
	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := loadFromFile(fileName)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}

	os.Remove(fileName)
}

func TestShuffle(t *testing.T) {
	deck := newDeck()
	deck2 := newDeck()

	deck.shuffle()

	if deck.toString() == deck2.toString() {
		t.Errorf("Expected deck to have a shuffled order")
	}
}

func TestDeal(t *testing.T) {
	deck := newDeck()
	dealt, left := deal(deck, 23)

	totalLen := len(dealt) + len(left)
	if totalLen != 52 {
		t.Errorf("Expected deck length of 52, but got %d", totalLen)
	}
}
