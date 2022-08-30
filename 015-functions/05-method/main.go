package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent {
		person: person {
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	sa1.intro()
	sa1.person.intro()
}

func (s secretAgent) intro() {
	if s.ltk {
		fmt.Printf("I'm %s %s. I have license to kill\n", s.person.first, s.person.last)
	} else {
		fmt.Printf("I'm %s, %s. I don't have license to kill\n", s.first, s.last)
	}
}

func (p person) intro() {
	fmt.Printf("I'm %s %s\n", p.first, p.last)
}
