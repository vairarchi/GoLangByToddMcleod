package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From even channel:", v)
		case v := <-o:
			fmt.Println("From odd channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok:", i, ok)
			} else {
				fmt.Println("from comma ok", i)
			}
			return
		}
	}
}
