/*
Callback
passing a func as an argument
functional programming not something that is recommended in go, however, it is good to be aware of callbacks
idiomatic go: write clear, simple, readable code
*/

package main

import "fmt"

func main() {
	fmt.Println("Sum of alll numbers ")
	ii := []int{5, 8, 9, 5, 6, 5, 8, 5, 6}
	s := sum(ii...)
	fmt.Println(s)
	fmt.Println("Sum of Even numbers by callback - functional programming")

	sumEven := even(sum, ii...)
	fmt.Println(sumEven)

	fmt.Println("Sum of Odd numbers by callback - functional programming")
	sumOdd := odd(sum, ii...)
	fmt.Println(sumOdd)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, xi ...int) int {
	vi := []int{}
	for _, v := range xi {
		if v%2 == 0 {
			vi = append(vi, v)
		}
	}
	return f(vi...)
}

func odd(f func(x ...int) int, xi ...int) int {
	vi := []int{}
	for _, v := range xi {
		if v%2 != 0 {
			vi = append(vi, v)
		}
	}
	return f(vi...)
}
