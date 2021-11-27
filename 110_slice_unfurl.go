/*
Unfurling a slice
When you have a slice of some type,
you can pass in the individual values in a slice by using the “...” operator.

code:
unfurling a slice of int
https://play.golang.org/p/T-mm6-SH71
passing in zero or more values
https://play.golang.org/p/Qc5sq_AK_T
variadic parameter has to be the final parameter
https://play.golang.org/p/euQ8PDQ8RN
*/
package main

import "fmt"

func main() {
	s := []int{1, 5, 8, 7, 6, 4, 6}
	// unfurling slice to match the argument type
	sum := sum(s...)
	fmt.Println(sum)
}

func sum(s ...int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
