/*
Variadic parameter
You can create a func which takes an unlimited number of arguments.
When you do this, this is known as a “variadic parameter.”
When use the lexical element operator “...T” to signify a variadic parameter
(there “T” represents some type).
code:
https://play.golang.org/p/Yi0FsQ2tKq
*/
package main

import "fmt"

func main() {
	sum := foo(5, 8, 6, 4, 8, 4)
	fmt.Println(sum)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
