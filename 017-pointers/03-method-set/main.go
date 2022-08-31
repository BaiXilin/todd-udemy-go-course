package main

import (
	"fmt"
	"math"
)

// method sets are methods attached to a type
// the receiver can be NON-POINTER or POINTER type

/*

receiver		value
--------------------------------
(t T)			T and *T
(t *T)			T and *T (T is addressable)
(t *T)			*T	 (T is unaddressable)

*/

type circle struct {
	radius float64
}

// Non-pointer receiver
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Pointer receiver
func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var size, length float64

	c := circle {
		radius: 5,
	}
	c_addr := &c

	// non-pointer value and non-pointer receiver
	size = c.area()
	fmt.Println("1. area =", size)

	// passing *circle to non-pointer receiver
	size = c_addr.area()
	fmt.Println("2. area =", size)

	// passing circle to pointer receiver
	length = c.perimeter()
	fmt.Println("3. perimeter =", length)

	// passing *circle to pointer receiver
	length = c_addr.perimeter()
	fmt.Println("4. perimeter =", length)
}


