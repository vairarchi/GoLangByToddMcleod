/*
Hands-on exercise #4
Fix the race condition you created in the previous exercise by using a mutex
it makes sense to remove runtime.Gosched()
code: https://github.com/GoesToEleven/go-programming
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	counter := 0
	gc := 50
	wg.Add(gc)
	for i := 0; i < 50; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Final Goroutines:", runtime.NumGoroutine())
}
