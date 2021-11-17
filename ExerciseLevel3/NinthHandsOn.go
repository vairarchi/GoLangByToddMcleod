/*

Hands-on exercise #9
Create a program that uses a switch statement with the switch expression
specified as a variable of TYPE string with the IDENTIFIER “favSport”.

solution: https://play.golang.org/p/h-8FnjpzWB

*/

package main

import "fmt"

func main() {
	favSport := "Git"
	// favSport := "Linkedin"
	switch favSport {
	case "Git":
		fmt.Println("Github: https://github.com/vairarchi/")
	case "Linkedin":
		fmt.Println("Linkedin: https://twitter.com/vairarchi/")
	}
}
