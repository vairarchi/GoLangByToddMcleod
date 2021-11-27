/*
Hands-on exercise #10
Closure is when we have “enclosed” the scope of a variable in some code block.
For this hands-on exercise, create a func which “encloses” the scope of a variable:
code: https://play.golang.org/p/a56uWtoFSL
*/

package main

import "fmt"

func main() {
	f1 := counter()
	f2 := counter()
	f1()
	f1()
	f2()
	f1()
	f2()
	f1()
	f2()
	f2()
	f1()
	f2()
	f1()
	f2()
	f1()
	f2()

}

func counter() func() {
	x := 0
	return func() {
		x++
		fmt.Println("Counter:", x)
	}
}
