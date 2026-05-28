package notifier

import "fmt"

type Sender interface {
	Send(message string) error
}

type ConsoleSender struct{}

func (ConsoleSender) Send(message string) error {
	fmt.Printf("console sender: %s\n", message)
	return nil
}

type ReminderService struct {
	sender Sender
}

func NewReminderService(sender Sender) ReminderService {
	return ReminderService{sender: sender}
}

func (s ReminderService) Remind(topic string) error {
	return s.sender.Send("practice reminder: " + topic)
}
