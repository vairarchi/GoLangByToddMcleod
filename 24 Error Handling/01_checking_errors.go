package main

import "fmt"

func main() {
	i, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
