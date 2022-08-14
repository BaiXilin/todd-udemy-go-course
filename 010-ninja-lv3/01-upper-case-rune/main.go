// Print every rune code point of the uppercase alphabet three times.

package main

import "fmt"

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
