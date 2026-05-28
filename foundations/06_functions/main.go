package main

import "fmt"

func square(n int) int {
	return n * n
}

func greet(name string) string {
	return "Welcome back, " + name + "."
}

func main() {
	fmt.Println("Lesson 06: Functions")
	fmt.Println(greet("Plato"))
	fmt.Printf("The square of %d is %d.\n", 6, square(6))
}
