package main

import "fmt"

func main() {
	fmt.Println("Lesson 10: Slices and Maps")

	topics := []string{"variables", "loops", "structs"}
	topics = append(topics, "interfaces")

	for i, topic := range topics {
		fmt.Printf("%d -> %s\n", i, topic)
	}

	progress := map[string]int{
		"Monday":    2,
		"Tuesday":   1,
		"Wednesday": 3,
	}

	for day, count := range progress {
		fmt.Printf("%s: %d lessons\n", day, count)
	}
}
