/*Hands-on exercise #9
A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
code: https://play.golang.org/p/0yGYPKh1y7
*/
package main

import "fmt"

func main() {
	values := []int{5, 8, 9, 5, 9, 4, 6, 4, 6, 5, 7, 2, 6, 4, 8, 6, 1}
	callbackFunc(sum, values...)

}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func callbackFunc(f func(x ...int) int, y ...int) {
	even := []int{}
	odd := []int{}
	for _, v := range y {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	fmt.Println("Sum of even: ", f(even...))
	fmt.Println("Sum of Odd:", f(odd...))
}
