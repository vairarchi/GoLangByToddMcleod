/*
Hands-on exercise #7
write a program that
launches 10 goroutines
each goroutine adds 10 numbers to a channel
pull the numbers off the channel and print them
solutions:
https://play.golang.org/p/R-zqsKS03P
https://play.golang.org/p/quWnlwzs2z
student solutions:
https://twitter.com/mannion
https://gist.github.com/mannion007/3c8899913974c1027ef6f13ec37b2b3f
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	const gorouines = 10
	// wg.Add(gorouines)
	c := make(chan int)
	for i := 0; i < gorouines; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- i
			}
			// wg.Done()
		}()
	}
	// wg.Wait()
	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}

	fmt.Println("Exiting......")
}
