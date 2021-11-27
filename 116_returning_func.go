/*
Returning a func
You can return a func from a func. Here is what that looks like.
code:
returning a string
https://play.golang.org/p/w3S9B1Vtyx
returning a func
step 1: https://play.golang.org/p/3Xk0wLFre8
step 2 - cleaned up: https://play.golang.org/p/NocmYezUP2
step 3 - cleaned up: https://play.golang.org/p/FSjvOfY0wW
step 4 - cleaned up: https://play.golang.org/p/7wbv9KNlhK
step 5 - cleaned up: https://play.golang.org/p/vW0IGeIAox
*/

package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())

	g := func() int {
		return 451
	}()

	fmt.Println(g)
}

func foo() func() int {
	return func() int {
		return 654
	}
}
