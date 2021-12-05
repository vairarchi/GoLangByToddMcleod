package main

import "fmt"

func main() {
	f()
	fmt.Println("returned normally from f()")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()
	fmt.Println("calling g")
	g(0)
	fmt.Println("returned normally from g()")
}

func g(s int) {
	if s > 3 {
		fmt.Println("Panicking....!")
		panic(fmt.Sprintf("%v", s))
	}
	defer fmt.Println("defer in g:", s)
	fmt.Println("Printing in g:", s)
	g(s + 1)
}
