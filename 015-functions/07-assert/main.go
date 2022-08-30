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

type human interface {
        speak()
}

func main() {
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

        enter_bar(sa1)
        enter_bar(p1)
}

func enter_bar(h human) {
	// using switch to determine exact type and act accordingly
	switch h.(type) {
		case person:
			// ASSERTING h is type person to access its members
			fmt.Printf("%s %s passed into a bar. He will have a beer.\n", h.(person).first, h.(person).last)
		case secretAgent:
			// ASSERTING h is type secretAgent to access its members
			fmt.Printf("%s %s passed into a bar. He can't drink because he is on mission.\n", h.(secretAgent).first, h.(secretAgent).last)
	}
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
