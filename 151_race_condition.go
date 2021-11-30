/*
Race condition

Here is a picture of the race condition we are going to create:
https://drive.google.com/file/d/0B22KXlqHz6ZNaW9mQ0U1b0tiUUE/view?usp=sharing&resourcekey=0--t6b5DJQck168b9R-BlQcw

Race conditions are not good code.
A race condition will give unpredictable results.
We will see how to fix this race condition in the next video.
code: https://play.golang.org/p/FYGoflKQej
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
