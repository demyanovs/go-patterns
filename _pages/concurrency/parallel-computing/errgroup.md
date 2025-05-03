---
layout: default
title: Error Group
description: "Handle multiple concurrent Go routines and propagate errors cleanly using Go’s errgroup package."
nav_order: 6
parent: Parallel Computing
grand_parent: Concurrency Patterns
permalink: /parallel-computing/errgroup
---

# Error Group

**Error Group** is a concurrency coordination pattern in Go, often used for parallel task execution with error management.

Useful when there is a large task that can be split into several subtasks.

There are two ways to use `errgroup`:
1. Using the WithContext method, which allows to pass the context to the group.
2. Without using the WithContext method.

In the first case, if one of the goroutines fails, all other goroutines will be cancelled.
In the second case, if one of the goroutines fails, all other goroutines will continue to run.

The following example demonstrates the use of both methods.

## Applicability
**Parallel Execution of Independent Tasks**.
When you need to run multiple tasks concurrently and wait for all of them to complete.

**Early Cancellation on Failure**.
When one failing task should cancel all other ongoing tasks—especially useful with errgroup.WithContext.

**Error Aggregation and Propagation**.
When you want to collect the first error encountered from a group of goroutines and return it.

**Concurrent I/O or Network Calls**.
Ideal for making multiple API calls, database queries, or file reads concurrently.

**Fan-out/Fan-in Workflows**.
When a task fans out into multiple subtasks that need to be gathered back (fan-in) after completion.

**Improving Performance Through Concurrency**.
When subtasks can be done in parallel to reduce overall execution time.

**Graceful Shutdown of Goroutines**.
When managing lifecycles of concurrent operations that should terminate cleanly if one fails.

**Simplified Goroutine Management**.
When you want to avoid manually tracking goroutines and their errors, and instead use structured concurrency.

**Context-Aware Concurrency**.
When the tasks should respect cancellation signals or timeouts via context.Context.

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

var errFailure = errors.New("some error")

func main() {
	ctx := context.Background()
	err := FetchUserDataWithError(ctx)
	// err := FetchUserDataWithoutError(ctx)

	if err != nil {
		fmt.Println("Error fetching user data:", err)
	}

	fmt.Println("Done")
}

func FetchUserDataWithError(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	// Run the first periodic task.
	group.Go(func() error {
		firstTask(ctx)
		return nil
	})

	// Run the second task that returns an error.
	group.Go(func() error {
		return secondTask()
	})

	// Wait for all goroutines to finish and return the first error (if any).
	return group.Wait()
}

func FetchUserDataWithoutError(ctx context.Context) error {
	var group errgroup.Group

	// Run the third periodic task.
	group.Go(func() error {
		thirdTask(ctx)
		return nil
	})

	// Run the fourth task that logs an error but doesn't return it.
	group.Go(func() error {
		fourthTask()
		return nil
	})

	// Wait for all goroutines to finish.
	return group.Wait()
}

func firstTask(ctx context.Context) {
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("first task running")
			if counter > 10 {
				return
			}
			counter++
		}
	}
}

func secondTask() error {
	fmt.Println("second task started")
	time.Sleep(3 * time.Second)
	fmt.Println("second task log error:", errFailure)
	return errFailure
}

func thirdTask(ctx context.Context) {
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("third task running")
			if counter > 10 {
				fmt.Println("third task finished")
				return
			}
			counter++
		}
	}
}

func fourthTask() {
	fmt.Println("fourth task started")
	time.Sleep(3 * time.Second)
	fmt.Println("fourth task log error:", errFailure)
}
```