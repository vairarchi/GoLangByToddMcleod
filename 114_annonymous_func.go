// Annonymous func - A self executing func
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello from Anonymous")
	}()

	foo := func(x int) {
		fmt.Println("Hello from anon taking argument: ", x)
	}

	foo(65)
}
