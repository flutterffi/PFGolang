package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/lockx"
)

func main() {
	fmt.Println("Lesson 53: Distributed Lock Shape")

	lock := lockx.New()
	fmt.Printf("first acquire: %v\n", lock.Acquire("maintenance"))
	fmt.Printf("second acquire: %v\n", lock.Acquire("maintenance"))
	lock.Release("maintenance")
	fmt.Printf("after release acquire: %v\n", lock.Acquire("maintenance"))
}
