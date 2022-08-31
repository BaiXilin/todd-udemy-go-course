package main

import "fmt"

func main() {
	x := 42

	fmt.Println("Before:")
	fmt.Printf("\tx = %d\n", x)
	fmt.Printf("\t&x = %p\n\n", &x)

	// pass by value
	take_value(x)
	fmt.Println("After take_value()")
	fmt.Printf("\tx = %d\n", x)
	fmt.Printf("\t&x = %p\n\n", &x)

	// pass by value (but value is pointer
	take_pointer(&x)
	fmt.Println("After take_pointer()")
	fmt.Printf("\tx = %d\n", x)
	fmt.Printf("\t&x = %p\n", &x)
}

func take_value(y int) {
	y = 0
	fmt.Println("Inside take_value()")
	fmt.Printf("\ty = %d\n", y)
	fmt.Printf("\t&y = %p\n\n", &y)
}

func take_pointer(y *int) {
	*y = 43
	fmt.Println("Inside take_pointer()")
        fmt.Printf("\ty = %d\n", y)
        fmt.Printf("\t&y = %p\n\n", &y)
}
