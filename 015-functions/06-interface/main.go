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

// interface defines a set of behavior
type human interface {
	speak()	// if you can speak, you are a human
}

func main() {
	// a value can have more than 1 type
	// here sa1 is type secretAgent.
	// but secretAgent also has speak method, so it is also of human interface
	sa1 := secretAgent {
		person: person {
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	p1 := person {
		first: "Dr.",
		last: "Yes",
	}

	// polymorphism: self_intro function takes in values of different type (actually, both are human type)
	self_intro(sa1)
	self_intro(p1)
}

func self_intro(h human) {
	fmt.Println("I am a human. I will introduce myself.")
	h.speak()
}

func (s secretAgent) speak() {
	if s.ltk {
		fmt.Printf("I'm %s %s. I have license to kill.\n", s.first, s.last)
	} else {
		fmt.Printf("I'm %s %s. I don't have license to kill.\n", s.first, s.last)
	}
}

func (p person) speak() {
	fmt.Printf("I'm %s %s.\n", p.first, p.last)
}

