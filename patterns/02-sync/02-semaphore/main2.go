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

const (
	maxGoroutines = 3
	totalJobs     = 10
)

func process(id int) Result {
	fmt.Printf("[%s]: running task %d\n", time.Now().Format("15:04:05"), id)
	time.Sleep(time.Second)

	return Result{
		Port:  id,
		State: true,
	}
}

func main() {
	var wg sync.WaitGroup

	results := make([]Result, 0, totalJobs)
	resultChan := make(chan Result, totalJobs)

	// Semaphore to limit concurrent goroutines
	semaphore := make(chan struct{}, maxGoroutines)

	for i := 1; i <= totalJobs; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() {
				<-semaphore // Release semaphore
			}()

			result := process(id)
			resultChan <- result
		}(i)
	}

	// Close the result channel once all tasks are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	for res := range resultChan {
		if res.State {
			results = append(results, res)
		}
	}

	fmt.Println("Results:", results)
}
