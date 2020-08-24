package main

import "fmt"

func main() {

	// A slice of string
	cards := []string{newCard(), "Ace of Spades"}
	cards = append(cards, "10 of hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	//	fmt.Println(cards)
}

func newCard() string {
	return "Five of diamonds"
}
