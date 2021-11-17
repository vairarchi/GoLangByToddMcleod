/*
	Hands-on exercise #1
	Write a program that prints a number in decimal, binary, and hex

	solution: https://play.golang.org/p/bAQxcEuK8O
*/

package main

import "fmt"

func main() {
	var a int = 56

	fmt.Printf("%d %b %#x", a, a, a)
}
