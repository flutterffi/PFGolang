package ratelimitx

import "time"

type Limiter struct {
	tokens chan struct{}
}

func New(rate int, interval time.Duration) *Limiter {
	limiter := &Limiter{
		tokens: make(chan struct{}, rate),
	}

	for i := 0; i < rate; i++ {
		limiter.tokens <- struct{}{}
	}

	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			for len(limiter.tokens) < cap(limiter.tokens) {
				limiter.tokens <- struct{}{}
			}
		}
	}()

	return limiter
}

func (l *Limiter) Allow() bool {
	select {
	case <-l.tokens:
		return true
	default:
		return false
	}
}
