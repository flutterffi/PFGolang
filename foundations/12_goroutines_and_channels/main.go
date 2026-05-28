package main

import "fmt"

func announce(task string, done chan<- string) {
	done <- "finished: " + task
}

func main() {
	fmt.Println("Lesson 12: Goroutines and Channels")

	done := make(chan string)

	go announce("read the syntax example", done)
	go announce("edit the code", done)

	fmt.Println(<-done)
	fmt.Println(<-done)
}
