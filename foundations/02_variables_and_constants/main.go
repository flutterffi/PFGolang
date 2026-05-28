package main

import "fmt"

func main() {
	const language = "Go"
	var learner string = "Plato"
	level := 1

	fmt.Println("Lesson 02: Variables and Constants")
	fmt.Printf("%s is learning %s at level %d.\n", learner, language, level)

	level++
	fmt.Printf("After one more session, the level becomes %d.\n", level)
}
