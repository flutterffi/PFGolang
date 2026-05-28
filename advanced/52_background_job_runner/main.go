package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/jobx"
)

func main() {
	fmt.Println("Lesson 52: Background Job Runner")

	runner := jobx.NewRunner()
	runner.Enqueue("rebuild search index")
	runner.Enqueue("send summary email")

	for _, result := range runner.RunAll() {
		fmt.Println(result)
	}
}
