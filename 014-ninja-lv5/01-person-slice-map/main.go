// Create your own type “person” which will have an underlying type of “struct” so that it can store
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavors.

// Take the code from the previous exercise, then store the values of type person in a map with the
// key of last name. Access each value in the map. Print out the values, ranging over the slice.

package main

import "fmt"

type person struct {
	first, last string
	flavor      string
}

func main() {
	p1 := person{
		first:  "James",
		last:   "Bond",
		flavor: "vanilla",
	}

	p2 := person{
		first:  "Miss",
		last:   "Moneypenny",
		flavor: "chocolate",
	}

	person_slice := []person{p1, p2}
	for _, p := range person_slice {
		fmt.Println(p.flavor)
	}

	person_map := map[int]person{
		1: p1,
		2: p2,
	}
	for _, v := range person_map {
		fmt.Println(v.flavor)
	}
}
