package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Err check 1:", ctx.Err())
	fmt.Println("No. Gorouines:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working:", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("Err check 2:", ctx.Err())
	fmt.Println("No. Gorouines:", runtime.NumGoroutine())

	fmt.Println("about to cancel")
	cancel()
	fmt.Println("cancelled context")
	fmt.Println("Err check 3:", ctx.Err())
	fmt.Println("No. Gorouines:", runtime.NumGoroutine())
}
