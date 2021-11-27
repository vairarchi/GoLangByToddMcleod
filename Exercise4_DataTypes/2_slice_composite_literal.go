/*
Hands-on exercise #2

Using a COMPOSITE LITERAL:
create a SLICE of TYPE int
assign 10 VALUES
Range over the slice and print the values out.
Using format printing
print out the TYPE of the slice

solution: https://play.golang.org/p/sAQeFB7DIs

*/

package main

import "fmt"

func main() {
	sl := []int{1, 5, 4, 8, 4, 6, 4, 6, 4, 6}
	for _, v := range sl {
		fmt.Printf("%v   -   %T \n", v, sl)
	}

}
