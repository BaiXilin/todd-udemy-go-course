// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
// slice:
// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."
// Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {
	s := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "MoneyPenny", "Hello James"},
	}

	for _, s0 := range s {
		for _, s1 := range s0 {
			fmt.Printf("%s\t", s1)
		}
		fmt.Println()
	}
}
