/*
Hands-on exercise #5
Create a variable of type string using a raw string literal. Print it.
solution: https://play.golang.org/p/dLy36A-V-w
*/

package main

import "fmt"

func main() {
	a := `here is something
	as a
	raw
	string literal
	
	"we can use double quotes"`

	fmt.Println(a)

}
