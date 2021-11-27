/*
Interfaces & polymorphism
In Go, values can be of more than one type.
An interface allows a value to be of more than one type.
We create an interface using this syntax:
“keyword identifier type” so for an interface it would be: “type human interface”
We then define which method(s) a type must have to implement that interface.
If a TYPE has the required methods,
which could be none (the empty interface denoted by interface{}),
then that TYPE implicitly implements the interface and is also of that interface type.
In Go, values can be of more than one type.
code: https://play.golang.org/p/rZH2Efbpot
*/

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secret agent speak")
}

// Using type switch
func foo(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Person was passed into foo - ", h.(person).first, h.(person).last)
	case secretAgent:
		fmt.Println("Secretagent was passed into foo - ", h.(secretAgent).first, h.(secretAgent).last, h.(secretAgent).ltk)
	}
	// fmt.Println("I called human", h)
}

func main() {
	p1 := person{
		first: "Todd",
		last:  "Mcleod",
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p1.speak()
	sa1.speak()

	foo(p1)
	foo(sa1)

}
