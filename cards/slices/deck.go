package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

// https://golang.org/pkg/strings/#Join
// a function named toString() that takes in a slice of string as the receiver
// and returns a string type output
func (d deck) toString() string {

	// This <return strings.Join(d, ",")> statement also worked
	return strings.Join([]string(d), ",")
}

func (d deck) saveToDrive(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromDisk(filename string) deck {

	// https://golang.org/pkg/io/ioutil/#ReadFile

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return s
}
