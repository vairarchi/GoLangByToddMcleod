package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 2
	b = 3

	fmt.Println(a)
	fmt.Printf("A is of type %T", a)

	fmt.Println(b)
	fmt.Printf("B is of type %T", b)

}
