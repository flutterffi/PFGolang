package main

import "fmt"

func main() {
	score := 82

	fmt.Println("Lesson 04: Conditionals")

	if score >= 90 {
		fmt.Println("Excellent work.")
	} else if score >= 75 {
		fmt.Println("Solid progress.")
	} else {
		fmt.Println("Keep practicing.")
	}

	if score%2 == 0 {
		fmt.Println("The score is even.")
	} else {
		fmt.Println("The score is odd.")
	}
}
