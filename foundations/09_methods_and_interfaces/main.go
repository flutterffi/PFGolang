package main

import "fmt"

type shape interface {
	area() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func printArea(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
}

func main() {
	fmt.Println("Lesson 09: Methods and Interfaces")
	printArea(rectangle{width: 4, height: 6})
}
