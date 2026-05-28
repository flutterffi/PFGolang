package main

import "fmt"

func main() {
	fmt.Println("Lesson 05: Loops")

	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
		fmt.Printf("i=%d sum=%d\n", i, sum)
	}

	words := []string{"focus", "practice", "repeat"}
	for index, word := range words {
		fmt.Printf("words[%d] = %s\n", index, word)
	}
}
