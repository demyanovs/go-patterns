package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Port  int
	State bool
}

const MAX_GOUROUTINE = 3
const TOTAL_JOBS = 10

func main() {
	var results []Result
	var wg sync.WaitGroup

	// Create a buffered channel to collect results
	resultChan := make(chan Result, TOTAL_JOBS)

	// Create a semaphore to limit concurrent goroutines
	// This prevents us from opening too many connections at once
	semaphore := make(chan struct{}, MAX_GOUROUTINE) // Limit to 100 concurrent scans

	// Launch goroutines for each port
	for i := 1; i <= TOTAL_JOBS; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }() // Release semaphore

			result := process2(i)
			resultChan <- result
		}(i)
	}

	// Close channel when all goroutines complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results from channel
	for result := range resultChan {
		if result.State {
			results = append(results, result)
		}
	}

	fmt.Println("Results:", results)
}

func process2(id int) Result {
	fmt.Printf("[%s]: running task %d\n", time.Now().Format("15:04:05"), id)
	time.Sleep(time.Second)

	return Result{
		Port:  id,
		State: true,
	}
}
