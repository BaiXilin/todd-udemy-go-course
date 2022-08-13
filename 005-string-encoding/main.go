package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Printf("is a %T\n\n", s)

	// converting string to byte slices
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("is a %T\n\n", bs)

	// print the bytes in different encoding
	for i, v := range bs {
		fmt.Printf("At location %d we have %#U ,or %#x or %#b\n", i, v, v, v)
	}
}
