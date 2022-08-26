package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Printf("p1 is of type %T\n", p1)

	fmt.Printf("%s %s is %d years old\n", p1.first, p1.last, p1.age)
}
