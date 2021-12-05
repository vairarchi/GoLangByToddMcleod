package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()

	_, err := os.Open("no-such-file.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func foo() {
	fmt.Println("after os.Exit() is called by fatal, deffered func foo() does not run")
}
