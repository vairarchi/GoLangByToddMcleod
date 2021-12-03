package main

import "fmt"

func main() {
	// Does not run because CHANNEL BLOCKS!!
	// doesNotRun()

	runsByGoRoutine()

	runsByBuffer()
}

func doesNotRun() {
	ch := make(chan int)
	ch <- 54
	fmt.Println(<-ch)
	// Does not run because CHANNEL BLOCKS!!
}

func runsByGoRoutine() {
	ch := make(chan int)
	go func() {
		ch <- 54
	}()
	fmt.Println(<-ch)
}

func runsByBuffer() {
	ch := make(chan int, 1)
	ch <- 58
	fmt.Println(<-ch)
}
