// Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
// TYPE []string which stores their favorite things. Store three records in your map. Print out all of
// the values, along with their index position in the slice.
// `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
// `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
// `no_dr`, `Being evil`, `Ice cream`, `Sunsets`
//
// Print out the map using range
//
// delete a record and print again

package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m {
		fmt.Printf("%s's favorite things are %v\n", k, v)
	}

	// delete no_dr from the map
	delete(m, `no_dr`)

	fmt.Println("After deletion")
	for k, v := range m {
		fmt.Printf("%s's favorite things are %v\n", k, v)
	}
}
