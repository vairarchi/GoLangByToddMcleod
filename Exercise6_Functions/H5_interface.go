/*
Hands-on exercise #5
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
code: https://play.golang.org/p/NGGikHNvHv
*/
package main

import "fmt"

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s square) area() {
	fmt.Println("Area of square is", s.length*s.length)
}
func (c circle) area() {
	fmt.Println("Area of circle is", c.radius*c.radius*3.14)
}

type shape interface {
	area()
}

func info(s shape) {
	s.area()
}

func main() {
	s1 := square{
		length: 12.5,
	}
	c1 := circle{
		radius: 6.8,
	}
	info(s1)
	info(c1)
}
