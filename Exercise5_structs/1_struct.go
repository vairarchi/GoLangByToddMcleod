/*
Hands-on exercise #1
Create your own type “person” which will have an underlying type of “struct”
so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values,
ranging over the elements in the slice which stores the favorite flavors.
solution:
https://play.golang.org/p/ouRHmH_POg

*/

package main

import "fmt"

type person struct {
	first        string
	last         string
	favIcecreams []string
}

func main() {
	p1 := person{
		first:        "james",
		last:         "bond",
		favIcecreams: []string{"vanilla", "butterscotch"},
	}

	p2 := person{
		first:        "todd",
		last:         "mcleod",
		favIcecreams: []string{"chocolate", "rajbhog"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for _, v := range p1.favIcecreams {
		fmt.Println(v)
	}

	fmt.Println(p2)

}
