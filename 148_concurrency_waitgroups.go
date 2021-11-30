/*
WaitGroup
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished. Writing concurrent code is super easy: all we do is put “go” in front of a function or method call.
runtime.NumCPU()
runtime.NumGoroutine()
sync.WaitGroup
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()
race conditions picture:
https://drive.google.com/file/d/0B22KXlqHz6ZNaW9mQ0U1b0tiUUE/view?usp=sharing&resourcekey=0--t6b5DJQck168b9R-BlQcw
code:
starting code:
https://play.golang.org/p/bnI0akWF9f
finished:
https://play.golang.org/p/VDioqpZ65h
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("No. of CPU:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("No. of CPU:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
