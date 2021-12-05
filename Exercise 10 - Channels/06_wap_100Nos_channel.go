/*
Hands-on exercise #6
write a program that
puts 100 numbers to a channel
pull the numbers off the channel and print them
solution: https://play.golang.org/p/096Lk1BR7o
*/
package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)
	for v := range c {
		fmt.Println(v)
	}
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
