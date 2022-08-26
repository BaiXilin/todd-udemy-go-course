package main

import "fmt"

func main() {
	// add to a map
	m := make(map[string]int)
	m["Mike"] = 80
	m["Eason"] = 90
	m["Alison"] = 100

	// ranging over key-value pairs in a map
	for k, v := range m {
		fmt.Printf("key: %s\tvalue: %d\n", k, v)
	}

	// can also ranging over index-value in a slice/array
	xi := []int{1, 2, 3, 4, 5}
	for i, v := range xi {
		fmt.Printf("xi[%d] = %d\n", i, v)
	}
}
