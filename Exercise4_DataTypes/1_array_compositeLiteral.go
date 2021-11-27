/*
	Hands-on exercise #1

	Using a COMPOSITE LITERAL:
	create an ARRAY which holds 5 VALUES of TYPE int
	assign VALUES to each index position.
	Range over the array and print the values out.
	Using format printing
	print out the TYPE of the array

	solution: https://play.golang.org/p/tD0SzV3hdf
*/

package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 5, 6, 3}
	for i, v := range arr {
		fmt.Println(i, " : ", v)
	}
}
