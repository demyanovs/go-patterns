---
layout: default
title: Mutex (Rate Limiting)
description: "Implement basic rate limiting and safe shared access in Go using Mutex locks."
nav_order: 1
parent: Synchronization
grand_parent: Concurrency Patterns
permalink: /sync/mutex
---

# Mutex (Rate Limiting)

The **Mutex** pattern is used to ensure that only one goroutine can access a shared resource at a time, 
preventing data races and ensuring thread safety. 

In Go, this can be implemented using channels as semaphores, where the mutex allows goroutines to "lock" a critical section and "unlock" it once done.

## Applicability

 - **Shared Resource Protection**.
Ensures only one goroutine accesses a shared resource at a time, preventing race conditions.

 - **Critical Section Management**.
Serializes access to a block of code, ensuring safe execution in concurrent environments.

 - **Custom Locking**.
Implements lightweight or custom locks using channels, without relying on sync.Mutex.

 - **Concurrency Control**.
Manages high concurrent access to shared resources, avoiding performance issues.

 - **Deadlock Prevention**.
Helps avoid deadlocks by managing resource access order in complex systems.

## Example

```go
package main

import (
	"fmt"
	"sync"
)

type Mutex struct {
	s chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{
		s: make(chan struct{}, 1),
	}
}

func (m *Mutex) Lock() {
	m.s <- struct{}{}
}

func (m *Mutex) Unlock() {
	<-m.s
}

const numGoroutines = 1000

func main() {
	m := NewMutex()
	counter := 0
	var wg sync.WaitGroup

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Mutex counter: %d\n", counter)
}
```
