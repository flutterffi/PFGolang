package main

import "fmt"

func JoinWords(words []string) string {
	result := ""
	for i, word := range words {
		if i > 0 {
			result += ","
		}
		result += word
	}

	return result
}

func main() {
	fmt.Println("Lesson 26: Benchmarks")
	fmt.Println(JoinWords([]string{"go", "bench", "practice"}))
	fmt.Println("Run: go test -bench=. ./advanced/26_benchmarks")
}
