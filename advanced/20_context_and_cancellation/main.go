package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func simulateWork(ctx context.Context, name string, delay time.Duration) error {
	select {
	case <-time.After(delay):
		fmt.Printf("%s completed in %s\n", name, delay)
		return nil
	case <-ctx.Done():
		return fmt.Errorf("%s stopped: %w", name, ctx.Err())
	}
}

func main() {
	fmt.Println("Lesson 20: Context and Cancellation")

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Millisecond)
	defer cancel()

	if err := simulateWork(ctx, "quick task", 60*time.Millisecond); err != nil {
		fmt.Println(err)
	}

	if err := simulateWork(ctx, "slow task", 200*time.Millisecond); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("The slow task hit the deadline.")
		}
		fmt.Println(err)
	}
}
