package main

import "fmt"

func main() {

	//	var card string = "Ace of Spades"
	// := is used only during initialization
	card := "Ace of Spades"
	newCard := newCard()
	fmt.Println(newCard)
	fmt.Println(card)
}

/*
function with return statement inside its body, should always
the type of value being returned.
*/
func newCard() string {
	return "Five of Diamonds"
}
