package main

import "fmt"

func main() {
	m := map[string]int{
		"james":   38,
		"giannis": 27,
	}

	m["vaibhav"] = 25
	for k, v := range m {
		fmt.Println(k, v)
	}
}
