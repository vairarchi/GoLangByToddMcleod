/*
Hands-on exercise #3
Create a struct “customErr” which implements the builtin.error interface.
Create a func “foo” that has a value of type error as a parameter.
Create a value of type “customErr” and pass it into “foo”.
If you need a hint, here is one.
solution:
https://play.golang.org/p/ixeowY2fd2
assertion
https://play.golang.org/p/pbl2kCYsM0
conversion
https://play.golang.org/p/1ldiBdkdzA
*/

package main

import "fmt"

type errCustom struct {
	info string
}

func (e errCustom) Error() string {
	return fmt.Sprintf("Here's the error: %v", e.info)
}

func main() {
	ec := errCustom{
		info: "Zero Error",
	}
	foo(ec)
	fmt.Println()
}

func foo(e error) {
	fmt.Println("foo ran -", e)
}
