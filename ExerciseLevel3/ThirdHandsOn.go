/*
	Hands-on exercise #4

	Create a for loop using this syntax
	for  { }
	Have it print out the years you have been alive.

	solution: https://play.golang.org/p/tnyqBPJ-i5
*/

package main

import "fmt"

func main() {
	i := 1996
	for {
		fmt.Println(i)
		if i >= 2021 {
			break
		}
		i++
	}
}
