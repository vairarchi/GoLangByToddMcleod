/*
Closure
one scope enclosing other scopes
variables declared in the outer scope are accessible in inner scopes
closure helps us limit the scope of variables
code:
scope of x: https://play.golang.org/p/YWuniJtu2R
scope of x narrowed to func main: https://play.golang.org/p/4hqrzybcFc
code block in code block with y: https://play.golang.org/p/6Hyqe_aU-R
enclosing a variable in a block of code: https://play.golang.org/p/fHez3lg8wc
*/

package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println("i of a:", a())
	fmt.Println("i of b:", b())
	fmt.Println("i of a:", a())
	fmt.Println("i of a:", a())
	fmt.Println("i of a:", a())
	fmt.Println("i of b:", b())
	fmt.Println("i of a:", a())
	fmt.Println("i of b:", b())
	fmt.Println("i of a:", a())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
