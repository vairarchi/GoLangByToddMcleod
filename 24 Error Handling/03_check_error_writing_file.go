package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("demo.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("Sample text to be sent to the newly created file...")
	io.Copy(f, r)
}
