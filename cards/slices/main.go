package main

import "fmt"

func main() {

	// To add more elements to a slice, use append.

	//cards := []string{"Ace of spades", newCard()}
	cards := newDeck()

	// Everytime you iterate over a range of slice, you throw the previous
	// values of iteration variables. Therefore, you need initialization
	// everytime.

	hand, remainingCards := deal(cards, 5)

	fmt.Println("Cards before dealing")
	cards.printCards()
	fmt.Println("Cards after dealing")
	hand.printCards()
	remainingCards.printCards()

	fmt.Println(cards.toString())

	fmt.Println(cards.saveToDrive("current_deck"))
}
