package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/calc"
)

func main() {
	fmt.Println("Lesson 15: Table-Driven Tests")
	fmt.Printf("Add(7, 5) = %d\n", calc.Add(7, 5))

	result, err := calc.Divide(20, 4)
	if err != nil {
		fmt.Printf("unexpected error: %v\n", err)
		return
	}

	fmt.Printf("Divide(20, 4) = %d\n", result)
	fmt.Println("Now open pkg/calc/calc_test.go and run: go test ./pkg/calc")
}
