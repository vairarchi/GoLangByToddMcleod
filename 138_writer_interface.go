/*
Writer interface
Understanding the writer interface from package io.
Also, one last note about working with JSON: encode & decode.
code: https://play.golang.org/p/3Txh-dKQBf

*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Fprintln(os.Stdout, "Hello")
	io.WriteString(os.Stdout, "Hello")
}
