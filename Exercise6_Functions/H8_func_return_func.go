/*
Hands-on exercise #8
Create a func which returns a func
assign the returned func to a variable
call the returned func
code: https://play.golang.org/p/c2HwqVE1Rd
*/
package main

import "fmt"

func main() {

	f := foo()
	v := f()
	fmt.Println(v)

}

func foo() func() int {
	x := 1
	fmt.Println("Inside called func 1")
	return func() int {
		x++
		return x
	}
}
