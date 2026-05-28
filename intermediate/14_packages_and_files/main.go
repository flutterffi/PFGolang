package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/tracker"
)

func main() {
	fmt.Println("Lesson 14: Packages and Multiple Files")

	session := tracker.NewSession("Read about packages", 30)
	fmt.Println(tracker.Summary(session))

	tracker.MarkComplete(&session)
	fmt.Println(tracker.Summary(session))
}
