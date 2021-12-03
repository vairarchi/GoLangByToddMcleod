/*
Fan in
Taking values from many channels, and putting those values onto one channel.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}
}

func send(even, odd chan<- int) {
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive(odd, even <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
