---
layout: default
title: Drop
description: "Discard excess data when buffers are full using the Drop Pattern to maintain system responsiveness under load."
nav_order: 3
parent: Stability Patterns
permalink: /stability/drop
---

# Drop

The **Drop** pattern is used to protect a system from overload by discarding incoming data when the processing queue or buffer is full. Instead of blocking or waiting, new input is immediately dropped, allowing the system to continue operating under high load without being overwhelmed. This pattern is especially useful in time-sensitive systems where stale data is less valuable than responsiveness.

It is typically implemented using a buffered channel with a select and default clause. If the channel is full, the default case handles the drop logic.

## Difference from other patterns
 - [Retry]({{ site.baseurl }}/stability/drop): attempt again later
 - [Timeout]({{ site.baseurl }}/stability/timeout): give up after a duration
 - [Circuit Breaker]({{ site.baseurl }}/stability/circuit-breaker): stop requests temporarily
 - **Drop**: give up immediately if the system is overloaded

## Example

```go
func drop() {
	const cap = 100
	ch := make(chan string, capacity)

	go func() {
		for p := range ch {
			fmt.Println("child: received signal:", p)
		}
	}()

	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- fmt.Sprintf("data %d", w):
			fmt.Println("parent: sent signal:", w)
		default:
			fmt.Println("parent: dropped data due to full buffer:", w)
		}
	}

	close(ch)
	fmt.Println("parent: sent shutdown signal")

	time.Sleep(time.Second)
}
```