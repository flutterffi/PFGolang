package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/notifier"
)

func main() {
	fmt.Println("Lesson 25: Interfaces In Design")

	service := notifier.NewReminderService(notifier.ConsoleSender{})
	if err := service.Remind("review interfaces and implementations"); err != nil {
		fmt.Printf("send error: %v\n", err)
	}

	fmt.Println("Swap ConsoleSender with a fake sender later to test behavior.")
}
