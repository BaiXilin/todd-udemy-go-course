package main

import "fmt"

type person struct {
	first, last string
	age int
}

func changeMe(p *person) {
	// the person grows one year older
	p.age++
}

func (p person) intro() {
	fmt.Printf("I'm %s %s. I am %d years old.\n", p.first, p.last, p.age)
}

func main() {
	p := person {
		first: "James",
		last: "Bond",
		age: 32,
	}
	p.intro()

	changeMe(&p)

	p.intro()
}
