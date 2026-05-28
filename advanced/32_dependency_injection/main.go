package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/notifier"
)

type fakeSender struct {
	messages []string
}

func (f *fakeSender) Send(message string) error {
	f.messages = append(f.messages, message)
	return nil
}

func main() {
	fmt.Println("Lesson 32: Dependency Injection")

	fake := &fakeSender{}
	service := notifier.NewReminderService(fake)
	if err := service.Remind("practice dependency injection"); err != nil {
		fmt.Printf("service error: %v\n", err)
		return
	}

	fmt.Printf("captured messages: %v\n", fake.messages)
	fmt.Println("The service depends on an interface, not a concrete sender.")
}
