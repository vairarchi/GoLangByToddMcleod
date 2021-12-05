/*
Printing & logging
You have a few options to choose from when it comes to printing out, or logging, an error message:
	fmt.Println()
	log.Println()
	log.Fatalln()
		os.Exit()
	log.Panicln()
		deferred functions run
		can use “recover”
panic()
code:
https://github.com/GoesToEleven/go-programming */

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-such-file.txt")
	if err != nil {
		fmt.Println("Error happened:", err)
		log.Println("Error happened:", err)
		log.Panic(err)
	}
	fmt.Println("Exiting...")
}
