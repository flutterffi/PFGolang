package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/queuex"
)

func main() {
	fmt.Println("Lesson 45: Message Queue Shape")

	publisher := &queuex.MemoryPublisher{}
	_ = publisher.Publish("tasks.created", "practice queue concepts")
	_ = publisher.Publish("tasks.completed", "1")

	fmt.Printf("published messages: %v\n", publisher.Messages)
}
