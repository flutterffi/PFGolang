package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/logx"
)

func main() {
	fmt.Println("Lesson 33: Structured Logging")

	if err := logx.Print("info", "starting practice step", map[string]any{
		"lesson": "33_structured_logging",
		"user":   "Plato",
	}); err != nil {
		fmt.Printf("log error: %v\n", err)
		return
	}

	if err := logx.Print("error", "sample failure for debugging", map[string]any{
		"component": "parser",
		"retryable": false,
	}); err != nil {
		fmt.Printf("log error: %v\n", err)
	}
}
