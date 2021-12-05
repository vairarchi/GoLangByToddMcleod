package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	l, err := os.Create("logs.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	log.SetOutput(l)
	f, err := os.Open("no-such-file.txt")
	defer f.Close()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("check logs.log for logs")
}
