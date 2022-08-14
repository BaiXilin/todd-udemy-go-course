// Create a program that uses a switch statement with the switch expression specified as a
// variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	favSport := "swimming"

	switch favSport {
	case "football":
		fmt.Println("Nah")
	case "basketball":
		fmt.Println("Nah")
	case "swimming":
		fmt.Println("Hell Yeah!")
	default:
		fmt.Println("Nerd, go out and exercise")
	}
}
