package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	tri := triangle{
		base:   4,
		height: 2,
	}

	sq := square{
		sideLength: 4,
	}

	printArea(tri)
	printArea(sq)
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Printf("The area is %.6f \n", area)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
