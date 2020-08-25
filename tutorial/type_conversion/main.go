package main

import "fmt"

// The need for type conversion from string to byte slice comes from, using
// WriteFile function of https://golang.org/pkg/io/ioutil/ package.

func main() {

	greetings := "Hi there!"
	fmt.Println([]byte(greetings))

}
