/*
Recursion
	- A function calling itself
*/
package main

import "fmt"

func main() {
	f := factorialRecursion(4)
	fmt.Println(f)
	f = factorialLoop(4)
	fmt.Println(f)
}
func factorialLoop(l int) int {
	prod := 1
	for ; l > 0; l-- {
		prod *= l

	}
	return prod
}
func factorialRecursion(l int) int {
	if l == 0 {
		return 1
	}
	return l * factorialRecursion(l-1)
}

// 4 * 3 * 2 * 1 * 1 * 1 * 1 * 1
