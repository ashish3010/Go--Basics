package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) circumf() float64 {
	return s.length * 4
}

func (c circle) circumf() float64 {
	return math.Pi * c.radius * 2
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func interfaceExample() {
	shapes := []shape{
		square{length: 5},
		circle{radius: 5},
		circle{radius: 3},
		square{length: 5},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("-----")
	}
}
