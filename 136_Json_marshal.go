/*
JSON marshal
Here is an example of how you marshal data in Go to JSON.
Also important, this video shows how the case of an identifier - lowercase or uppercase, determines whether or not the data can be exported.
code:
https://play.golang.org/p/jtE_n16cQO
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "james",
		Last:  "bond",
		Age:   45,
	}
	p2 := Person{
		First: "money",
		Last:  "penny",
		Age:   45,
	}
	people := []Person{p1, p2}
	byteData, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteData))
}
