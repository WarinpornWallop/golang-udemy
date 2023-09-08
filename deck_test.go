package main

import (
	"os"
	"testing"
)

// func TestNewDeck(t *testing.T) {: This line defines a new test function called TestNewDeck
// that takes a single argument, t, of type *testing.T. In Go, test functions must start with the word "Test,"
// and they are used to run various test cases.
func TestNewDeck(t *testing.T) {
	d := newDeck() // Call the newDeck function to create a new deck of cards.
	// Check if the length of the deck 'd' is equal to 16. If not, it's an error.
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	// Check if the first card in the deck 'd' is "Ace of Spades." If not, it's an error.
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	// Check if the last card in the deck 'd' is "Four of Clubs." If not, it's an error.
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
// This test is focused on two functions: saveToFile and newDeckFromFile. 
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Cleanup: Remove any existing "_decktesting" file to start with a clean slate.
    os.Remove("_decktesting")

    // Create a new deck of cards.
    deck := newDeck()

    // Save the newly created deck to a file named "_decktesting".
    deck.saveToFile("_decktesting")

    // Load a deck from the saved file "_decktesting".
    loadedDeck := newDeckFromFile("_decktesting")

    // Check if the loaded deck has 16 cards as expected.
    if len(loadedDeck) != 16 {
        // If the loaded deck doesn't have the expected length, report an error.
        t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
    }

    // Cleanup: Remove the "_decktesting" file after testing.
    os.Remove("_decktesting")
}