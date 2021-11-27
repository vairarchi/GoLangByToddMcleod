/*
Hands-on exercise #2

create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in

create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in

code: https://play.golang.org/p/B0yRxtBQPD
*/

package main

import "fmt"

func main() {
	x := []int{48, 8, 9, 5, 4, 7, 6, 5, 6, 1, 5, 8, 2, 4, 5, 2, 54, 15, 1, 5}
	fmt.Println("foo: ", foo(x...))
	fmt.Println("bar: ", bar(x))
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
