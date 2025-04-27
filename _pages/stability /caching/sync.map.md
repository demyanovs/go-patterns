---
layout: default
title: Caching with sync.Map
description: "Simplify concurrent cache access in Go by using the built-in sync.Map for thread-safe operations."
nav_order: 2
parent: Caching Patterns
grand_parent: Stability Patterns
permalink: /stability/caching/with-sync-map
---

# Caching with sync.Map

The `sync.Map` type in Go provides a concurrent map implementation optimized for scenarios with many goroutines accessing and modifying the map. 
It is particularly useful for caching, where reads dominate writes. 
Unlike a regular map with manual locking (`sync.Mutex`), `sync.Map` internally manages concurrency with minimal contention, making it more efficient for read-heavy workloads.

Typical use cases for caching with `sync.Map` include memoization, lazy loading, and storing frequently accessed computed results. 
It supports safe concurrent `Load`, `Store`, `LoadOrStore`, and `Delete` operations without needing explicit locks.

## Benefits of `sync.Map` over `map` + `sync.RWMutex`
 - **Built-in Concurrency Optimizations**<br/>
`sync.Map` internally reduces lock contention, especially for read-heavy scenarios.

 - **Simpler Code for Concurrency**<br/>
No need to manage `Lock`, `Unlock`, `RLock`, or `RUnlock` manually - fewer chances for deadlocks or forgetting to release locks.

 - **Ready-Made Utilities**<br/>
Methods like `LoadOrStore`, `LoadAndDelete`, and `Range` simplify common concurrent operations.

 - **Better for Many Goroutines**<br/>
Handles extremely high concurrent access more efficiently than a `map` + `RWMutex` if reads dominate writes.

## Drawbacks of `sync.Map` compared to `map` + `sync.RWMutex`
 - **No Type Safety**<br/>
Keys and values are `interface{}` - you must do type assertions manually (`value.(YourType)`).
This can lead to runtime panics and harder-to-maintain code.

- **Worse Performance on Write-Heavy Workloads**<br/>
  When there are lots of writes and deletes, `sync.Map`'s internal mechanics become less efficient than a simple `RWMutex`-guarded `map`.

- **More Boilerplate and Casting**<br/>
Extra lines for type checking and casting clutter your code.

- **Harder to Refactor**<br/>
Changing the type of values later is risky because the compiler won't catch mistakes.

- **Higher Memory Overhead**<br/>
`sync.Map` maintains separate structures for "read" and "dirty" states, consuming more memory compared to a plain `map`.

- **Less Flexibility**<br/>
Fine-grained operations (like conditional updates, batch updates) are harder or impossible to implement cleanly.

## Example

```go
package main

import (
  "fmt"
  "sync"
)

type Cache struct {
  data sync.Map
}

func (c *Cache) Get(key string) (string, bool) {
  val, ok := c.data.Load(key)
  if !ok {
    return "", false
  }
  return val.(string), true
}

func (c *Cache) Set(key, value string) {
  c.data.Store(key, value)
}

func main() {
  cache := &Cache{}

  cache.Set("language", "Go")

  if val, found := cache.Get("language"); found {
    fmt.Println("Found:", val)
  } else {
    fmt.Println("Not found")
  }
}
```