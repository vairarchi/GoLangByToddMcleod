/*
We can slice a slice which means that we can cut parts of the slice away.
We do this with the colon operator.
*/

package main

import "fmt"

func main() {
	x := []int{1, 4, 45, 97, 5, 8, 85}

	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}
