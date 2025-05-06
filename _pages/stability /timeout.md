---
layout: default
title: Timeout
description: "Manage timeouts in Go to ensure responsiveness and prevent stuck operations."
nav_order: 2
parent: Stability Patterns
permalink: /stability/timeout
---

# Timeout

The **timeout** pattern is used to **limit the amount of time an operation can take**. 
If the operation doesn't complete within the specified time, it is **aborted** or handled appropriately.

This pattern is particularly useful for network requests, database queries, or any long-running processes that must not block indefinitely.

## Example 1: Using `context.WithTimeout`

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second): // Simulates long work
		return nil
	case <-ctx.Done(): // Context timeout/cancellation
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := longRunningOperation(ctx)
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation completed successfully")
	}
}
```

## Example 2: Using `select` with `time.After`

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	resultChan := make(chan string)

	go func() {
		time.Sleep(3 * time.Second) // Simulate work
		resultChan <- "Success"
	}()

	select {
	case result := <-resultChan:
		fmt.Println("Received:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! Operation took too long.")
	}
}
```