/*Hands-on exercise #3
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
code: https://play.golang.org/p/XluEuUD0Nw
*/
package main

import "fmt"

func foo() {
	fmt.Print("Foo called first but defer used")
}

func bar() {
	fmt.Println("Bar called seciond without defer")
}

func main() {
	defer foo()
	bar()
}
