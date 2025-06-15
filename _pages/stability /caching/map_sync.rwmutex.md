---
layout: default
title: Caching with Map and sync.RWMutex
description: "Build concurrent-safe caches in Go using standard maps protected by sync.RWMutex."
nav_order: 1
parent: Caching Patterns
grand_parent: Stability Patterns
permalink: /stability/caching/with-map-and-rwmutex
---

# Caching with Map and sync.RWMutex
A simple and efficient caching technique in Go where a `map` is used to **store key-value** pairs, 
and a `sync.RWMutex` protects **concurrent access**. `RLock` is used for **reads** to allow multiple readers simultaneously, 
and `Lock` is used for **writes** to ensure exclusive access when modifying the cache.

## Example

```go
package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, found := c.data[key]
	return val, found
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func main() {
	cache := NewCache()

	cache.Set("username", "gopher")

	if val, found := cache.Get("username"); found {
		fmt.Println("Found:", val)
	} else {
		fmt.Println("Not found")
	}
}
```