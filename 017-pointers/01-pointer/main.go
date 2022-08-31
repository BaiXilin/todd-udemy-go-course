package main

import "fmt"

func main() {
	value := 42
	addr := &value	// & operator gives the address

	fmt.Println("value =", value)
	fmt.Printf("type of value = %T\n", value)
	fmt.Println("addr =", addr)
	fmt.Printf("type of addr = %T\n\n", addr)
	
	// changing the value through address changes the value of the original variable
	addr_copy := addr
	*addr_copy = 43 // * operator gives the value stored at the address (dereference)
	fmt.Printf("value = %d\n", *&value)
}
