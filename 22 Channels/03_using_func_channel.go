package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)

	// send
	go bar(ch)

	// receive
	foo(ch)
}

// send
func bar(c chan<- int) {
	c <- 3
}

//receive
func foo(c <-chan int) {
	fmt.Println(<-c)
}
