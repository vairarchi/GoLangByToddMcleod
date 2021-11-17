/*
A multi-dimensional slice is like a spreadsheet.
You can have a slice of a slice of some type.

Code:
https://play.golang.org/p/S4cyB89Zpm
*/

package main

import "fmt"

func main() {
	x := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println(x)

	y := []string{"miss", "powerpuff", "girls", "vodka"}
	fmt.Println(y)

	xy := [][]string{x, y}
	fmt.Println(xy)
}
