package main

import "fmt"

func main() {
	x := []int{98, 78, 65, 21}
	fmt.Println(x)

	x = append(x, 32, 98, 78, 65)
	fmt.Println(x)

	y := []int{999, 929, 949, 959}
	x = append(x, y...)
	fmt.Println(x)

	// Using builtin append() to delete element 32 from slice
	x = append(x[:4], x[5:]...)
	fmt.Println(x)

}
