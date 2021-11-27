/*
Hands-on exercise #4
Create and use an anonymous struct.
*/
package main

import "fmt"

func main() {
	p1 := struct {
		name   string
		age    int
		weight float64
	}{
		name:   "Vaibhav",
		age:    27,
		weight: 96.5,
	}

	fmt.Println(p1)
}
