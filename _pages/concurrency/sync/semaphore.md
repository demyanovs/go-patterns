---
layout: default
title: Semaphore (Rate Limiting)
description: "Control resource access and rate limiting in Go with the Semaphore synchronization primitive."
nav_order: 2
parent: Synchronization
grand_parent: Concurrency Patterns
permalink: /sync/semaphore
---

# Semaphore (Rate Limiting)
The **Semaphore** pattern is used to **limit** the number of **concurrently running goroutines**. 

It helps manage resource usage, such as limiting the number of simultaneous network requests or file operations. 
This is typically implemented using a buffered channel to act as a counting semaphore - each goroutine must acquire a slot before proceeding and release it when done.

```mermaid
graph LR
    start["Start Tasks"] --> acquire["Acquire Semaphore (buffered channel)"]
    acquire --> worker["Do Work"]
    worker --> release["Release Semaphore"]
    release --> done["Done"]

```

## Applicability
 - **Limiting Concurrent Operations**.
Use when you need to restrict the number of concurrently running goroutines, such as database queries, HTTP requests, or file reads/writes.

- **Rate-Limiting External Services**.
  Prevents overwhelming external APIs or services that have concurrency or rate limits.

 - **Managing Resource-Intensive Tasks**. 
Ideal when tasks consume significant system resources (CPU, memory, network) and running too many in parallel would degrade performance.

 - **Avoiding System Overload**.
Helps ensure stability in applications that spawn many background jobs or workers by capping their concurrency.

- **Controlling Access to Shared Resources**.
Useful for limiting access to shared resources (e.g., open connections, memory pools) without full locking mechanisms.

## Example

```go
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
```
