/*
******************** bcrypt *********************************
All too often today you still hear about websites and databases
being hacked and user’s information being compromised.
There is no excuse for this.
As programmers, we have the tools to protect user data.
Bcrypt is one of the tools you can use to protect user data.
Using bcrypt, we will gain further understanding as to how to read
and implement code from the standard library.
code:
https://goo.gl/ZVKnRx
*/

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwordStored := `password12345`
	fmt.Println(passwordStored)
	bs, err := bcrypt.GenerateFromPassword([]byte(passwordStored), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	passwordEntered := `password12345`
	fmt.Println(passwordEntered)
	err = bcrypt.CompareHashAndPassword(bs, []byte(passwordEntered))
	if err != nil {
		fmt.Println("Wrong password entered")
	} else {
		fmt.Println("Password matched!")
	}

}
