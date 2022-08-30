// Create a func which returns a func
// ● assign the returned func to a variable
// ● call the returned func

package main

import "fmt"

func main() {
	x := meaning_of_life()
	fmt.Println("The meaning of life is", x())
}

func meaning_of_life() func() int {
	return func() int {
		return 42
	}
}
