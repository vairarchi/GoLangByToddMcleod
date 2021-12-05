package main

import "os"

func main() {
	_, err := os.Open("no-such-file.txt")
	if err != nil {
		panic(err)
	}
}
