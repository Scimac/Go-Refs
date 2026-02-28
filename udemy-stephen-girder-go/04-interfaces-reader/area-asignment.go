package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLenght float64
}

type shape interface {
	getArea() float64
}

func AreaAssigment() {
	fmt.Println("Area calculation using Interfaces")
	t1 := triangle{
		base:   10,
		height: 5,
	}

	s1 := square{
		sideLenght: 4,
	}

	fmt.Println("\nArea of Triangle:\n")
	printArea(t1)
	fmt.Println("\nArea of Square:\n")
	printArea(s1)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println("Area is: ", area)
}
