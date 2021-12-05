/*
Hands-on exercise #2
Get this code running:
https://play.golang.org/p/oB-p3KMiH6
solution: https://play.golang.org/p/isnJ8hMMKg
https://play.golang.org/p/_DBRueImEq
solution: https://play.golang.org/p/mgw750EPp4
*/
package main

import (
	"fmt"
)

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	fmt.Println(cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
