// Create a for loop using this syntax
// 	for {}  OR  for condition {}
// Have it print out the years you have been alive

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 2001; i <= time.Now().Year(); i++ {
		fmt.Println(i)
	}

	fmt.Println()

	j := 2001
	for {
		if j > time.Now().Year() {
			break
		}
		fmt.Println(j)
		j++
	}
}
