package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %s", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', but got %s", d[len(d)-1])
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	// clean up before and after the test
	// always name the temp file something unique to avoid conflicts
	os.Remove("_decktesting")

	d := newDeck()
	err := d.saveToFile("_decktesting")

	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}

	loadedDeck, err := loadFromFile("_decktesting")

	if err != nil {
		t.Errorf("Error loading deck from file: %v", err)
	}

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
