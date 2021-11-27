/*
Hands-on exercise #3
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
solution: https://play.golang.org/p/SGfiULXzAB

*/

package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[0:5])
	fmt.Println(x[5:10])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
