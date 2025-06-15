---
layout: default
title: Circuit Breaker
description: "Implement the Circuit Breaker pattern in Go to handle failures gracefully and keep your applications resilient and responsive."
nav_order: 4
parent: Stability Patterns
permalink: /stability/circuit-breaker
---

# Circuit Breaker

The **Circuit Breaker** pattern prevents an application from repeatedly trying to execute an operation that's likely to fail. 
It detects failures and stops further attempts for a period, allowing the system to recover or degrade gracefully.

## Applicability

 - Call external services (like APIs or databases) that might become unavailable.
 - Want to avoid cascading failures in distributed systems. 
 - Need to quickly fail operations instead of waiting for timeouts.

## Example

```go
package main

import (
	"errors"
	"fmt"
	"time"
)

type CircuitBreaker struct {
	failures    int
	state       string
	lastAttempt time.Time
}

func NewCircuitBreaker() *CircuitBreaker {
	return &CircuitBreaker{state: "CLOSED"}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	now := time.Now()
	if cb.state == "OPEN" && now.Sub(cb.lastAttempt) < 5*time.Second {
		return errors.New("circuit is open, request blocked")
	}

	err := fn()
	if err != nil {
		cb.failures++
		cb.lastAttempt = now
		if cb.failures >= 3 {
			cb.state = "OPEN"
		}
		return err
	}

	cb.reset()
	return nil
}

func (cb *CircuitBreaker) reset() {
	cb.failures = 0
	cb.state = "CLOSED"
}

func main() {
	cb := NewCircuitBreaker()
	service := func() error {
		return errors.New("service down")
	}

	for i := 0; i < 5; i++ {
		err := cb.Call(service)
		if err != nil {
			fmt.Println("Attempt", i+1, ":", err)
		}
		time.Sleep(1 * time.Second)
	}
}
```