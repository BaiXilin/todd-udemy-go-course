// closure: limiting the scope of a variable
// goal: the scope of variables should be as narrow as possible
package main

import "fmt"

func main() {
	// a, b holds two separate function values. Each with a different copy of x
	a := incrementor()
	b := incrementor()

	fmt.Println("a1:", a())
	fmt.Println("a2:", a())
	fmt.Println("a3:", a())

	fmt.Println("b1:", b())
	fmt.Println("b2:", b())
}

func incrementor() func() int {
	// the anonymous function has access to x, so the scope of x is extended into that anonymous function
	var x int
	return func() int {
		x++
		return x
	}
}
