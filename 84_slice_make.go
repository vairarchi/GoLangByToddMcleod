/*
	Slices are built on top of arrays. A slice is dynamic in that it will grow in size.
	The underlying array, however, does not grow in size.
	When we create a slice, we can use the built in function make to specify
	how large our slice should be and also how large the underlying array should be.
	This can enhance performance a little bit.

		make([]T, length, capacity)
		make([]int, 50, 100)

	code
	https://play.golang.org/p/07hH1b-hvD
*/

package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 45
	x[9] = 54
	// This will create an error -
	// panic: runtime error: index out of range [10] with length 10
	// x[10] = 99

	// We will use the builtin append function to increase the len of our slice
	x = append(x, 9874)

	// watch what happens to the length of the slice
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 98, 7)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
