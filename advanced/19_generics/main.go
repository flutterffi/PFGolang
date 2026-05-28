package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/genericset"
)

func max[T ~int | ~float64 | ~string](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println("Lesson 19: Generics")

	fmt.Printf("max(8, 5) = %d\n", max(8, 5))
	fmt.Printf("max(3.14, 6.28) = %.2f\n", max(3.14, 6.28))
	fmt.Printf("max(\"go\", \"gym\") = %s\n", max("go", "gym"))

	topics := genericset.New("loops", "structs", "tests", "loops")
	fmt.Printf("set has \"tests\": %v\n", topics.Has("tests"))
	fmt.Printf("unique topics: %v\n", topics.Values())
}
