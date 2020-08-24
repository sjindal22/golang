/*
Package is like a project or a workspace. It is a collection of common source
file. Package main, is used only when we want the script to spit out an
executable file. Therefore, you can run and test your code right away.
*/

package main

/*
fmt, abbreviated as format, is a standard go lang library. And by making
an import of the library, you are letting package main access all of its functions.
*/

/* In short, package main has no access to any of the GO libraries (pkgs) by default,
and you would need to make a connection using the import statment, for main to
access all code with in the library.
*/

import "fmt"

func main() {

	cards := "Ace of Spades"
	fmt.Println(cards)
	fmt.Println("Hi there!")
}
