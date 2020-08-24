package main

import "fmt"

// Create a new type deck which is a slice
// of strings but with some extended functionalities.

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Club", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// "d deck" is a receiver
// d is a variable attached to the type "deck". By convention only a first few
// characters of the type are used as a variable name for the association.
func (d deck) printCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
