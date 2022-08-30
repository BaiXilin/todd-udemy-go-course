package main

import "fmt"

func main() {
	// we can assign an anonymous function to a variable
	// function is a TYPE, like int, string...
	f := func() {
		fmt.Println("My first function expression")
	}
	f()
}
