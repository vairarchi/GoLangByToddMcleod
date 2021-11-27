/*
Hands-on exercise #2

Take the code from the previous exercise,
then store the values of type person in a map with the key of last name.
Access each value in the map.

Print out the values, ranging over the slice.
solution: https://play.golang.org/p/RvvLyAMvGo
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
		favIcecreams: []string{"white", "yellow"},
	}
	p2 := person{
		first:        "todd",
		last:         "mcleod",
		favIcecreams: []string{"vanilla", "chocochips"},
	}

	mapOfStructs := map[string]person{
		"bond":   p1,
		"mcleod": p2,
	}

	fmt.Println(mapOfStructs)
	for k, v := range mapOfStructs {
		fmt.Println(k, v.first, v.last)
		for _, r := range v.favIcecreams {
			fmt.Println("\t", r)
		}

	}
}
