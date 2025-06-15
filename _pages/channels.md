---
layout: default
title: Channels
description: "Learn how to use channels in Go for effective goroutine communication. Explore unbuffered and buffered channels with examples, best practices, and recommendations."
nav_order: 50
permalink: /articles/channels
nav_exclude: true
---

# Channels

Channels in Go are a powerful concurrency primitive that allow goroutines to communicate by sending and receiving values. Channels can be either **unbuffered** or **buffered**.

**Unbuffered** channels require both a sender and receiver to be ready at the same time. Sending blocks until another goroutine receives, and receiving blocks until a value is sent.

**Buffered** channels have a fixed capacity. Sending blocks only when the buffer is full, and receiving blocks only when the buffer is empty.

This table summarizes the behavior of channel operations (`close`, `send`, and `receive`), depending on the channel state:

| Operation | Nil Channel      | Closed Channel      | Open Channel                   |
|:----------|:-----------------|:--------------------|:-------------------------------|
| **Close** | Panic             | Panic                | Succeeds                       |
| **Send**  | Blocks forever    | Panics               | Blocks (if full) or sends       |
| **Receive** | Blocks forever  | Never blocks (returns zero value) | Blocks (if empty) or receives |

## Recommendations for Using Channels

- **Close channels only where they are created**.<br/>
  The responsibility for closing a channel should stay with the sender (creator), not receivers.

- **Receiving goroutines should never attempt to close a channel**.<br/>
  Closing from multiple places can cause panics ("close of closed channel") and is hard to coordinate safely.

## Example 1: Unbuffered Channel

```go
package main

import "fmt"

func main() {
    ch := make(chan int) // unbuffered channel

    // Sender goroutine
    go func() {
        fmt.Println("Sending 123...")
        ch <- 123 // blocks until another goroutine receives
        fmt.Println("Sent 123")
    }()

    // Receiver (in main goroutine)
    value := <-ch // blocks until a value is sent
    fmt.Println("Received:", value)
}
```

## Example 2: Buffered Channel

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 2) // buffered channel with capacity 2

    ch <- 1 // succeeds immediately (buffer has space)
    fmt.Println("Sent 1")

    ch <- 2 // succeeds immediately (buffer still has space)
    fmt.Println("Sent 2")

    // Next send would block because buffer is full
    // ch <- 3 // would block here if uncommented

    fmt.Println(<-ch) // receives 1
    fmt.Println(<-ch) // receives 2
}

```