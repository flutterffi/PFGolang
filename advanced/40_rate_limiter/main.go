package main

import (
	"fmt"
	"time"

	"github.com/flutterffi/PFGolang/pkg/ratelimitx"
)

func main() {
	fmt.Println("Lesson 40: Rate Limiter")

	limiter := ratelimitx.New(2, 80*time.Millisecond)
	for requestID := 1; requestID <= 5; requestID++ {
		allowed := limiter.Allow()
		fmt.Printf("request=%d allowed=%v\n", requestID, allowed)
		time.Sleep(35 * time.Millisecond)
	}
}
