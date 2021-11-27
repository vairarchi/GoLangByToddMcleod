/*
Methods
A method is nothing more than a FUNC attached to a TYPE. When you attach a func to a type it is a method of that type. You attach a func to a type with a RECEIVER.
code: https://play.golang.org/p/HIBt2HHiIm
*/
package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(params) (return/s) {code}
func (s secretAgent) speak() {
	fmt.Println("My name is ", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}
	sa1.speak()
}
