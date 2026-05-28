package main

import "fmt"

func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func main() {
	fmt.Println("Lesson 07: Multiple Returns")

	if result, ok := divide(18, 3); ok {
		fmt.Printf("18 / 3 = %.2f\n", result)
	}

	if _, ok := divide(18, 0); !ok {
		fmt.Println("Cannot divide by zero.")
	}
}
