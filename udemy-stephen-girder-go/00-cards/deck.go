package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// deck is a custom type that represents a slice of strings
// It is used to represent a deck of cards.
type deck []string

// NewDeck creates and returns a new deck of cards
func newDeck() deck {
	// Let's be smart about it
	cards := deck{}

	// suits of cards
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	// ranks of cards
	ranks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Create a new deck of cards
	for _, suit := range suits {
		for _, rank := range ranks {
			// Append the card to the deck
			// Using fmt.Sprintf to format the string
			cards = append(cards, fmt.Sprintf("%s of %s", rank, suit))
		}
	}

	// Return the deck of cards
	return cards
}

// print prints the contents of the deck
func (d deck) print() {
	for i, card := range d {
		// Print the index and the card
		// Using fmt.Printf for formatted output
		// %v is used to print the value of the card
		// %d is used to print the index
		fmt.Printf("%d: %v\n", i, card)
	}
}

// deal takes a deck and a hand size, and returns two decks:
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert the whole deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// saveToFile saves the deck to a file
func (d deck) saveToFile(filename string) error {
	cardsString := d.toString()

	// The third argument (0666) sets the file permissions.
	// Common permissions:
	// 0644 - Owner can read/write, others can read
	// 0600 - Owner can read/write, others have no permissions
	// 0666 - Everyone can read/write (not recommended for sensitive files)
	// 0755 - Owner can read/write/execute, others can read/execute (for executables)
	err := os.WriteFile(filename, []byte(cardsString), 0644)
	return err
}

// loadFromFile loads a deck from a file
func loadFromFile(filename string) (deck, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// Convert the byte slice to a string and split it into a slice of strings
	cards := strings.Split(string(data), ", ")
	return deck(cards), nil
}

// shuffle function
func (d deck) shuffle() {
	// Implementing a simple shuffle algorithm
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i] // Swap cards
	}
}

// shuffle with custom random seed
func (d deck) shuffleWithSeed() {
	// Implementing a simple shuffle algorithm with a custom seed
	// using current time in nanoseconds as a seed for better precision for randomness
	n := time.Now().UnixNano()

	seededRand := rand.New(rand.NewSource(n)) // Example seed

	for i := len(d) - 1; i > 0; i-- {
		j := seededRand.Intn(i + 1)
		d[i], d[j] = d[j], d[i] // Swap cards
	}
}
