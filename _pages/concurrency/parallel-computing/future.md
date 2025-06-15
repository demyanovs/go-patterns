---
layout: default
title: Future (Promise)
description: "Learn how Future (Promise) can be used in Go to handle asynchronous results."
nav_order: 5
parent: Parallel Computing
grand_parent: Concurrency Patterns
permalink: /parallel-computing/future
---

# Future (Promise)

The **Future** pattern allows starting a computation in the background and retrieve the result later, 
allowing other work to continue in the meantime.

```mermaid
graph LR
    start["Start Task"] --> initiate["Initiate Future (async)"]
    initiate --> task["Do Task (in goroutine)"]
    task --> future["Future (holds result)"]
    future --> result["Get Result (from future)"]
    result --> done["Done"]
```

## Applicability

 - **Deferred Computation**.
When you know you'll need a result later, but don't want to block the current execution.

 - **Concurrent I/O Operations**.
Useful for performing multiple network requests or disk reads in parallel.

 - **Parallel Task Execution**.
When you can execute independent tasks simultaneously to improve performance.

 - **UI or API Responsiveness**.
Helps avoid blocking the main thread or request handler while a background job is running.

 - **Lazy Initialization**.
Start expensive setup only when the result is eventually needed, without blocking early.

## Example

```go
package main

import (
	"fmt"
	"time"
)

type data struct {
	Body  string
	Error error
}

func main() {
	future1 := future("https://example1.com")
	future2 := future("https://example2.com")

	fmt.Println("Requests started")

	body1 := <-future1
	body2 := <-future2

	fmt.Printf("Response 1: %v\n", body1)
	fmt.Printf("Response 2: %v\n", body2)
}

func future(url string) <-chan data {
	resultChan := make(chan data, 1)

	go func() {
		body, err := doGet(url)

		resultChan <- data{Body: body, Error: err}
	}()

	return resultChan
}

func doGet(url string) (string, error) {
	time.Sleep(time.Millisecond * 200)
	return fmt.Sprintf("Response of %s", url), nil
}
```