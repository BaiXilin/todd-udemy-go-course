// Create a program that uses a switch statement with no switch expression specified.

package main

import "fmt"

func main() {
	// if a switch doesn't specify an expression or identifier, use true by default
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 3):
		fmt.Println("This should not print")
	case (42 == 42):
		fmt.Println("This should print")
	default:
		fmt.Println("This should not print")
	}
}
