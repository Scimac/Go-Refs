package main

import "fmt"

func main() {
	fmt.Println("Welcome to the cards game!")

	cards := newDeck()

	// Print the deck of cards
	// cards.print()

	// Deal 5 cards from the deck
	hand, remainingDeck := deal(cards, 5)

	fmt.Println("\nHand:")
	hand.print()

	fmt.Println("\nRemaining Deck:")
	remainingDeck.print()

	fmt.Println("\nDeck as a string:")
	fmt.Println(cards.toString())

	fmt.Println("\nSaving the deck to a file...")
	err := cards.saveToFile("my_cards.txt")
	if err != nil {
		fmt.Println("Error saving the deck to a file:", err)
	} else {
		fmt.Println("Deck saved successfully!")
	}

	fmt.Println("\nLoading the deck from a file...")
	loadedDeck, err := loadFromFile("my_cards.txt")
	if err != nil {
		fmt.Println("Error loading the deck from a file:", err)
	} else {
		fmt.Println("Deck loaded successfully!")
		loadedDeck.print()
	}
	fmt.Println("\nDeck length:", len(cards))

	fmt.Println("\nShuffling the deck...")
	cards.shuffleWithSeed()
	fmt.Println("\nShuffled Deck:")
	cards.print()
}
