package main

import "fmt"

func main() {
	m := map[string]int {
		"Mike": 80,
		"Eason": 90,
		"Alison": 100,
	}

	fmt.Println(m)

	// delete a key from the map
	delete(m, "Mike")

	// note that no error is thrown if deleting a non-existent key
	delete(m, "James")

	fmt.Println(m)
}
