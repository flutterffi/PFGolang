package metricsx

import "sync"

type Counter struct {
	mu     sync.Mutex
	values map[string]int
}

func NewCounter() *Counter {
	return &Counter{values: map[string]int{}}
}

func (c *Counter) Add(name string, delta int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[name] += delta
}

func (c *Counter) Snapshot() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()

	copyMap := make(map[string]int, len(c.values))
	for key, value := range c.values {
		copyMap[key] = value
	}

	return copyMap
}
