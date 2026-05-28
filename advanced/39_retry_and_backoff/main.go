package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/flutterffi/PFGolang/pkg/retryx"
)

func main() {
	fmt.Println("Lesson 39: Retry and Backoff")

	failures := 0
	err := retryx.Do(4, 25*time.Millisecond, func(attempt int) error {
		fmt.Printf("attempt %d\n", attempt)
		if failures < 2 {
			failures++
			return errors.New("temporary network issue")
		}

		fmt.Println("operation succeeded")
		return nil
	})
	if err != nil {
		fmt.Printf("final error: %v\n", err)
	}
}
