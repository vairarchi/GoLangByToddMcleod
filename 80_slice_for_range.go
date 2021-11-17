package main

import "fmt"

func main() {
	x := []int{5, 4, 7, 8, 9, 7, 65}
	fmt.Println(len(x))
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
}
