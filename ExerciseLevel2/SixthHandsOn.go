/*
Hands-on exercise #6
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
solution: https://play.golang.org/p/MDLF3v5EGT

*/

package main

import "fmt"

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
