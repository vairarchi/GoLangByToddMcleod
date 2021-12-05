/*
Hands-on exercise #4
Starting with this code, pull the values off the channel using a select statement
solution: https://play.golang.org/p/FulKBY5JNj
*/
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 0
		close(c)
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case v := <-q:
			fmt.Println("Closing", v)
			return
		}
	}
}
