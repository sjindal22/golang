package main

import "fmt"

func main() {

	//var car map[string]string

	car := make(map[string]string)

	car["audi"] = "A7"
	car["toyota"] = "corolla"

	fmt.Println(car)

	delete(car, "toyota")

	fmt.Println(car)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#f24b75",
	}

	fmt.Println(colors)

	printColors(colors)
}

func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex code for " + color + " is " + hex)
	}
}
