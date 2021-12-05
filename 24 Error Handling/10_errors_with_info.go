package main

import (
	"errors"
	"fmt"
	"log"
)

var userDefinedError = errors.New("vaibhav defined error: squareroot of negative number")

func main() {
	fmt.Printf("Type of userDefinedError: %T\n", userDefinedError)
	_, err := sqrt(-32)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, userDefinedError
	}
	return 10, nil
}
