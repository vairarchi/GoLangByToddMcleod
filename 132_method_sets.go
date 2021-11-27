/*
Method sets
Method sets determine what methods attach to a TYPE.
It is exactly what the name says: What is the set of methods for a given type?
That is its method set.

IMPORTANT: “The method set of a type determines the INTERFACES that the type implements.....”
~ golang spec
The above “important” was left out of this video
but will be discussed in the “concurrency” section in a video called “method sets revisited”.

a NON-POINTER RECEIVER
works with values that are POINTERS or NON-POINTERS.
a POINTER RECEIVER
only works with values that are POINTERS.

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

code:
NON-POINTER RECEIVER & NON-POINTER VALUE
https://play.golang.org/p/2ZU0QX12a8
NON-POINTER RECEIVER & POINTER VALUE
https://play.golang.org/p/glWZmm0gY6
POINTER RECEIVER & POINTER VALUE
https://play.golang.org/p/pWFxsg6MSe
POINTER RECEIVER & NON-POINTER VALUE
https://play.golang.org/p/G3lEy-4Mc8 ( code does not run )
this codes does run - notice the difference -  method set determines
the INTERFACES that the type implements
https://play.golang.org/p/KK8gjsAWBZ
*/
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	// info(c)
	fmt.Println(c.area())
}
