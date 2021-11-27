/**Hands-on exercise #6
Build and use an anonymous func
code: https://play.golang.org/p/DQX3xEIcRe  */
package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is a annonymous func")
	}()
	f := func() {
		fmt.Println("This is func expression")
	}
	f()
}
