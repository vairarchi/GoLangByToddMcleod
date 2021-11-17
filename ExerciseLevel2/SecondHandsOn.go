/*
	Hands-on exercise #2
	Using the following operators, write expressions and assign their values to variables:
	==
	<=
	>=
	!=
	<
	>

	Now print each of the variables.

	solution: https://play.golang.org/p/76R-poSzaY
*/

package main

import (
	"fmt"
)

func main() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)
}
