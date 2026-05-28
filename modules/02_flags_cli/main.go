package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Go learner", "the learner name to greet")
	repeat := flag.Int("repeat", 1, "how many times to print the greeting")
	flag.Parse()

	fmt.Println("Module 02: Flags CLI")
	for i := 0; i < *repeat; i++ {
		fmt.Printf("%d. Welcome, %s.\n", i+1, *name)
	}
}
