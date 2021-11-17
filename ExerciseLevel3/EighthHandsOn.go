/*
Hands-on exercise #8
Create a program that uses a switch statement with no switch expression specified.
solution: https://play.golang.org/p/YpPgkWGqKY
*/

package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("False")
	}
}
