/*
a numbered sequence of elements of a single type
does not change in size
used for Go internals; generally not recommended for your code
https://golang.org/ref/spec#Array_types
*/
package main

import "fmt"

func main() {
	var x [5]int
	x[3] = 42
	fmt.Printf("%v %T", x, x)
}
