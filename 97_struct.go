/*
Struct
A struct is an composite data type.
(composite data types, aka, aggregate data types, aka, complex data types).
Structs allow us to compose together values of different types.
code: https://play.golang.org/p/hNI_rSK-C6
*/

package main

import "fmt"

type Person struct {
	first string
	last  string
}

func main() {
	p1 := Person{
		first: "james",
		last:  "bond",
	}

	p2 := Person{
		first: "Money",
		last:  "Penny",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last)
}
