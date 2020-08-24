package main

import "fmt"

func main() {

	//	var card string = "Ace of Spades"
	// ":=" creates a new variable and is used only during initialization. It
	// infers the type of variable with the value being initialized to it.

	// Basic declaration and assignment of the variable
	card := "Ace of Spades"
	anotherCard := newCard()
	fmt.Println(anotherCard)
	fmt.Println(card)

	// Using slices
	// Arrays & Slice are two data structure, where arrays are of fixed size and slice
	// can grow or shrink in size. However, the type of variables must be same in both.
	// You read []string{} as the "slice of string"
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
