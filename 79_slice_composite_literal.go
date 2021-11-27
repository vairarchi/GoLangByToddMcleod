/*
	Built on top of an array
	holds value of the same type
	changes in size
	has a length and capacity
	https://golang.org/ref/spec#Slice_types
*/

// a SLICE allows you to group together VALUES of same data types

package main

import "fmt"

func main() {
	// x := type{vaues}   //Composite literal
	x := []int{4, 589, 8, 8, 9, 8, 95, 9, 4, 55}
	fmt.Println(x)
}
