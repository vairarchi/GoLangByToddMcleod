package main

import "fmt"

func main() {
	cb := make(chan int)
	cs := make(<-chan int)
	cr := make(chan<- int)
	fmt.Printf("cb\t%T\n", cb)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)
}
