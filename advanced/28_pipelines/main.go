package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/pipeline"
)

func main() {
	fmt.Println("Lesson 28: Pipelines")

	values := pipeline.Generate(1, 2, 3, 4, 5, 6)
	squared := pipeline.Square(values)
	evenSquares := pipeline.FilterEven(squared)

	for value := range evenSquares {
		fmt.Printf("pipeline output: %d\n", value)
	}
}
