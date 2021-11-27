/*
What are pointers?

All values are stored in memory.
Every location in memory has an address.
A pointer is a memory address.
&
*int
*
code
code from video: https://play.golang.org/p/Ysv5Adn3V1
step 1 - take an address & : https://play.golang.org/p/BO7zRQN4Xm
step 2 - dereference * : https://play.golang.org/p/ucJYPu3QkP
step 3 - dereference * : https://play.golang.org/p/KpLImTmQCa
*/
package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	fmt.Println(&a) // & gives us the address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	fmt.Printf("%T\n", *&a)
	fmt.Println()
	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives us the value stored at the particular address
	fmt.Println(*&a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", *b)
}
