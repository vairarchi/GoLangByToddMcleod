/*
Hands-on exercise #1
in addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
code: https://github.com/GoesToEleven/go-programming
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go foo()
	wg.Wait()
	wg.Add(1)
	go bar()
	wg.Wait()
}

func foo() {
	// runtime.Gosched()
	fmt.Println("Hello from foo :)")
	wg.Done()
}

func bar() {
	fmt.Println("hello from bar :(")
	wg.Done()
}
