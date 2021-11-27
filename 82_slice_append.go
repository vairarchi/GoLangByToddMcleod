/*
To append values to a slice, we use the special built in function append.
This function returns a slice of the same type.
*/

package main

import "fmt"

func main() {
	x := []int{45, 8, 9}
	fmt.Println(x)

	x = append(x, 98, 87, 65)
	fmt.Println(x)

	y := []int{999, 999, 999}
	// variadic
	x = append(x, y...)
	fmt.Println(x)
}
