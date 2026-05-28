package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	ID    int
	Value int
}

type result struct {
	Worker int
	JobID  int
	Output int
}

func worker(id int, jobs <-chan job, results chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(20 * time.Millisecond)
		results <- result{
			Worker: id,
			JobID:  job.ID,
			Output: job.Value * job.Value,
		}
	}
}

func main() {
	fmt.Println("Lesson 21: Worker Pool")

	jobs := make(chan job)
	results := make(chan result, 5)

	var wg sync.WaitGroup
	for workerID := 1; workerID <= 3; workerID++ {
		wg.Add(1)
		go worker(workerID, jobs, results, &wg)
	}

	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- job{ID: i, Value: i + 1}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("worker=%d job=%d output=%d\n", result.Worker, result.JobID, result.Output)
	}
}
