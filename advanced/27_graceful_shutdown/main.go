package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ctx context.Context, done chan<- struct{}) {
	ticker := time.NewTicker(40 * time.Millisecond)
	defer ticker.Stop()
	defer close(done)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker: received cancellation, cleaning up")
			time.Sleep(30 * time.Millisecond)
			fmt.Println("worker: cleanup complete")
			return
		case tick := <-ticker.C:
			fmt.Printf("worker: heartbeat at %s\n", tick.Format("15:04:05.000"))
		}
	}
}

func main() {
	fmt.Println("Lesson 27: Graceful Shutdown")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signals)

	done := make(chan struct{})
	go worker(ctx, done)

	go func() {
		time.Sleep(140 * time.Millisecond)
		signals <- syscall.SIGTERM
	}()

	sig := <-signals
	fmt.Printf("main: received signal %s\n", sig)
	cancel()
	<-done
	fmt.Println("main: shutdown complete")
}
