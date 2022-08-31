package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

// the circle returned here is unaddressable
func create_circle(r float64) circle {
	return circle{
		radius: r,
	}
}

func create_circle_ptr(r float64) *circle {
	return &circle{
		radius: r,
	}
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var size, length float64

	// unaddressable non-pointer type and non-pointer receiver method - no problem
	size = create_circle(5).area()
	fmt.Println("area =", size)

	// pointer type and pointer receiver method -- no problem
	length = create_circle_ptr(5).perimeter()
	fmt.Println("perimeter =", length)

	// unaddressable non-pointer type cannot be passed into pointer receiver method
	// size := create_circle(5).area()	// error
	// fmt.Println("size =", size)

	// unaddressable non-pointer type and non-pointer receiver method - no-problem
	size = create_circle(5).area()
	fmt.Println("area =", size)
}
