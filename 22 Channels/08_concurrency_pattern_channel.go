/*
Fan out
Taking some work and putting the chunks of work onto many goroutines.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanInOut(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("Exiting.....")
}

func populate(c1 chan<- int) {
	for i := 0; i < 1000; i++ {
		c1 <- i
	}
	close(c1)
}

func fanInOut(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return rand.Intn(1000)
}
