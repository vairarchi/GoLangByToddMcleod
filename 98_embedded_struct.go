/*
Embedded structs
We can take one struct and embed it in another struct.
When you do this the inner type gets promoted to the outer type.
code: https://play.golang.org/p/u6b3qTr1CH
*/

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
			age:   32,
		},
		licenseToKill: true,
	}

	p2 := person{
		first: "Money",
		last:  "Penny",
		age:   25,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.first)
}
