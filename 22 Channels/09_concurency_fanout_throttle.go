// Restricting the no. of go routine sot be used
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

	go fanOut(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Exiting..........")
}

func populate(c1 chan<- int) {
	defer close(c1)
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	// close(c1)
}

func fanOut(c1, c2 chan int) {
	defer close(c2)
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- functionThatTakesTime(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func functionThatTakesTime(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return rand.Intn(500)
}
