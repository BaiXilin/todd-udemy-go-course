// ● create a type SQUARE
// ● create a type CIRCLE
// ● attach a method to each that calculates AREA and returns it
// ○ circle area= π r 2
// ○ square area = L * W
// ● create a type SHAPE that defines an interface as anything that has the AREA method
// ● create a func INFO which takes type shape and then prints the area
// ● create a value of type square
// ● create a value of type circle
// ● use func info to print the area of square
// ● use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type shape interface {
	calc_area() float64
}

type square struct {
	length, width float64
}

type circle struct {
	radius float64
}

func (s square) calc_area() float64 {
	return s.length * s.width
}

func (c circle) calc_area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	area := s.calc_area()
	switch s.(type) {
	case square:
		fmt.Println("The area of this square is", area)
	case circle:
		fmt.Println("The area of this circle is", area)
	}
}

func main() {
	s := square{
		length: 5,
		width:  10,
	}

	c := circle{
		radius: 1,
	}

	info(s)
	info(c)
}
