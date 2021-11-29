package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 5, 2, 5, 8, 6, 3, 2, 4, 7, 8, 2, 1, 0, 5, 4, 0}
	xs := []string{"fs", "the", "very expensice", "to be very sure", "f", "gdsg", ""}
	fmt.Println("Before:", xi)
	sort.Ints(xi)
	fmt.Println("After: ", xi)
	fmt.Println("Before:", xs)
	sort.Strings(xs)
	fmt.Println("After: ", xs)

}
