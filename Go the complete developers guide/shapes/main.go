package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base, height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 5}

	printArea(t)
	printArea(s)
}

func (tri triangle) getArea() float64 {
	return tri.base * tri.height / 2
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
