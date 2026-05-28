package main

import "fmt"

type Lesson struct {
	Title    string
	Duration int
	Complete bool
}

func main() {
	fmt.Println("Lesson 08: Structs")

	lesson := Lesson{
		Title:    "Struct Basics",
		Duration: 25,
		Complete: false,
	}

	fmt.Printf("%+v\n", lesson)

	lesson.Complete = true
	fmt.Printf("Updated lesson status: %+v\n", lesson)
}
