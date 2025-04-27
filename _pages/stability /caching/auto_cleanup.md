---
layout: default
title: Caching with Automatic Cleanup
description: "Implement automatic expiration and cleanup for cached items in Go to manage memory usage effectively."
nav_order: 3
parent: Caching Patterns
grand_parent: Stability Patterns
permalink: /stability/caching/with-expiration
---

# Caching with Automatic Cleanup

This caching pattern stores data with a time-to-live (TTL) and automatically removes expired entries in the background. 
A separate goroutine periodically scans the cache and deletes outdated items, keeping memory usage efficient without requiring manual checks during `Get` operations.

## Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	Value      string
	Expiration int64
}

func (item Item) Expired() bool {
	return time.Now().UnixNano() > item.Expiration
}

type AutoCleanupCache struct {
	data    map[string]Item
	mu      sync.RWMutex
	ttl     time.Duration
	cleanup time.Duration
	stop    chan struct{}
}

func NewAutoCleanupCache(ttl, cleanupInterval time.Duration) *AutoCleanupCache {
	cache := &AutoCleanupCache{
		data:    make(map[string]Item),
		ttl:     ttl,
		cleanup: cleanupInterval,
		stop:    make(chan struct{}),
	}
	go cache.startCleanup()
	return cache
}

func (c *AutoCleanupCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = Item{
		Value:      value,
		Expiration: time.Now().Add(c.ttl).UnixNano(),
	}
}

func (c *AutoCleanupCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.data[key]
	if !found || item.Expired() {
		return "", false
	}
	return item.Value, true
}

func (c *AutoCleanupCache) startCleanup() {
	ticker := time.NewTicker(c.cleanup)
	for {
		select {
		case <-ticker.C:
			c.deleteExpired()
		case <-c.stop:
			ticker.Stop()
			return
		}
	}
}

func (c *AutoCleanupCache) deleteExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now().UnixNano()
	for key, item := range c.data {
		if item.Expiration <= now {
			delete(c.data, key)
		}
	}
}

func (c *AutoCleanupCache) Stop() {
	close(c.stop)
}

func main() {
	cache := NewAutoCleanupCache(2*time.Second, 1*time.Second)

	cache.Set("token", "xyz789")

	time.Sleep(1 * time.Second)

	if val, found := cache.Get("token"); found {
		fmt.Println("Found before expiration:", val)
	}

	time.Sleep(2 * time.Second)

	if val, found := cache.Get("token"); found {
		fmt.Println("Still found:", val)
	} else {
		fmt.Println("Expired and cleaned up!")
	}

	cache.Stop()
}
```