package main

import "fmt"

func main() {
	// declaration using literals
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}

	fmt.Println(m)

	// make declaration
	// n := make(map[string]int)

	// accessing a key & value
	fmt.Printf("James is %d years old\n", m["James"])

	// maybe a key doesn't exist in the map
	// comma ok idiom
	if v, ok := m["Mr.Banana"]; !ok {
		fmt.Println("Mr.Banana is a fantacy", v) // if a key doesn't exist, zero value will be returned
	}
}
