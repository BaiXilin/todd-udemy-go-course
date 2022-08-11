package main

import "fmt"

var x int = 42

type num_hotdogs int

func main() {
	var y num_hotdogs = 10

	fmt.Printf("x has value %v is of type %T\n", x, x)
	fmt.Printf("The number of hotdogs is %v of type %T\n", y, y)

	// cannot assign because x and y are two different types (have to use conversion)
	// x = y

	// conversion
	x = int(y)
	fmt.Printf("The number of hotdogs is %v of type %T after conversion\n", x, x)
}
