package main

import "fmt"

func main() {

	//	var card string = "Ace of Spades"
	// := is used only during initialization

	// Basic declaration and assignment of the variable
	card := "Ace of Spades"
	anotherCard := newCard()
	fmt.Println(anotherCard)
	fmt.Println(card)

	// Using slices
	multipleCards := []string{"Ace of dimaond", card, anotherCard}
	for i, card := range multipleCards {
		fmt.Println(i, card)
	}

}

/*
function with return statement inside its body, should always
the type of value being returned.
*/
func newCard() string {
	return "Five of Diamonds"
}
