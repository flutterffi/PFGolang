package queuex

type Publisher interface {
	Publish(topic, message string) error
}

type MemoryPublisher struct {
	Messages []string
}

func (p *MemoryPublisher) Publish(topic, message string) error {
	p.Messages = append(p.Messages, topic+":"+message)
	return nil
}
