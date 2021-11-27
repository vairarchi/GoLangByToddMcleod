package main

import "fmt"

func main() {
	foo()
	bar("foo")
	s := fooBar("fooBar")
	x, y := multipleReturn("arg1", "arg2")
	fmt.Println(s)
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(parameter string) {
	fmt.Println("Hello from ", parameter)
}

func fooBar(param1 string) string {
	s := fmt.Sprint("Hello from ", param1)
	return s
}

func multipleReturn(param1, param2 string) (string, int) {
	s := fmt.Sprint("Hello from ", param1, param2)
	c := len(s)

	return s, c
}
