package main

import "fmt"

func main() {
	x := book_burn()
	fmt.Printf("x is type %T\n", x)

	fmt.Println("Book burns at", x())
	fmt.Println("To confirm, book burns at", book_burn()())
}

// returns a function that returns an int
func book_burn() func() int {
	return func() int {
		return 451
	}
}
