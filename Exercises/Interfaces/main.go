package main

import (
	"fmt"
	"math"
)

// shap interface
// a contract that have a set of functions to be implemented
type shape interface {
	area() float64
	circumf() float64
}

// A struct is a collection of data fields with declared data types.
type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return math.Pi * 2 * c.radius
}

// the printShapInfo method getting the interface "shape"
// so we can pass multiple structs into the same function where we want the same behavior.
func printShapInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumferece of %T is: %0.2f \n", s, s.circumf())

}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		square{length: 151.2},
		circle{radius: 18.4},
	}
	for _, v := range shapes {
		printShapInfo(v)
		fmt.Println("--------")
	}

}
