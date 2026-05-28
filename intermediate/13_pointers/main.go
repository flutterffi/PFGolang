package main

import "fmt"

type Counter struct {
	Value int
}

func increment(counter *Counter, step int) {
	counter.Value += step
}

func reset(counter *Counter) {
	counter.Value = 0
}

func main() {
	fmt.Println("Lesson 13: Pointers")

	counter := Counter{Value: 10}
	fmt.Printf("Before increment: %+v\n", counter)

	increment(&counter, 5)
	fmt.Printf("After increment: %+v\n", counter)

	reset(&counter)
	fmt.Printf("After reset: %+v\n", counter)
}
