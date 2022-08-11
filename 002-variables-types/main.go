package main

import "fmt"

// declare a variable outside of func body using var keyword
var y = 42

// cannot use short declaration outside of func body
// x := 1

func main() {
	x := "shaken, but not stirred"
	fmt.Printf("%s is of type %T\n", x, x)

	fmt.Printf("%d is of type %T\n", y, y)

	// Golang is a static programming language
	// cannot assign other types of value to a variable
	// y = "this is a sentence"
}
