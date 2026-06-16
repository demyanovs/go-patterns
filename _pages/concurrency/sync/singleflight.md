---
layout: default
title: Singleflight
description: "Singleflight in Go: deduplicate concurrent identical calls so only one executes while others share its result."
nav_order: 3
parent: Synchronization
grand_parent: Concurrency Patterns
permalink: /sync/singleflight
---

# Singleflight

The **Singleflight** pattern ensures that when multiple goroutines concurrently request the same piece of work (identified by a key), only one of them actually executes it. The rest simply wait and receive a shared copy of the result.

This is commonly used to prevent a **cache stampede** (a.k.a. "thundering herd"): when a cache entry expires, many concurrent requests for the same key would otherwise all hit the database or a downstream service at the same time. Singleflight collapses them into a single call.

## Applicability

 - **Cache Stampede Prevention**.
When a cache entry expires, avoids dozens of goroutines simultaneously recomputing or refetching the same value.

 - **Deduplicating Expensive Calls**.
When concurrent requests for the same resource (e.g. the same user ID, the same file) should share one expensive call instead of each making their own.

 - **Reducing Load on Backends**.
Protects databases, APIs, or other downstream services from duplicate concurrent traffic for identical keys.

 - **Per-Key Coordination**.
Unlike a global lock, only requests for the *same key* are coalesced — requests for different keys still run concurrently.

 - **Concurrent Handler Calls for the Same Resource**.
When an HTTP handler is hit concurrently by many users requesting the same resource (e.g. `/users/42`), and the expensive lookup can't simply be moved out of the request path, singleflight ensures only one of those concurrent calls actually performs the lookup while the rest share its result.

## Example: Deduplicating Concurrent Calls

`singleflight.Group` provides a ready-made implementation. `Do` executes `fn` for a given key, and any calls that arrive while it is in flight wait for and share its result instead of calling `fn` again.

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	group     singleflight.Group
	callCount int
	mu        sync.Mutex
)

// fetchUser simulates an expensive call, e.g. a database query or HTTP request.
func fetchUser(userID string) (string, error) {
	mu.Lock()
	callCount++
	mu.Unlock()

	time.Sleep(200 * time.Millisecond) // simulate latency
	return fmt.Sprintf("user-data-for-%s", userID), nil
}

func main() {
	var wg sync.WaitGroup

	// 5 goroutines request the same user concurrently.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			v, err, shared := group.Do("user:42", func() (interface{}, error) {
				return fetchUser("42")
			})
			if err != nil {
				fmt.Println("error:", err)
				return
			}

			fmt.Printf("caller %d got %q (shared=%v)\n", n, v, shared)
		}(i)
	}

	wg.Wait()

	mu.Lock()
	fmt.Println("fetchUser was called", callCount, "time(s)")
	mu.Unlock()
}
```

Even though five goroutines call `Do` with the same key at the same time, `fetchUser` runs only **once**; the other four calls block until it completes and receive the same result, with `shared` set to `true`.

## Example: Singleflight + Redis

A common case is an HTTP handler that reads from a Redis cache. On a cache miss, every concurrent request for the same key would otherwise query the database (or another slow backend) and overwrite the cache independently. Wrapping the "check cache, fall back to DB, populate cache" path in `singleflight.Group.Do` ensures only one of those concurrent requests does the work, while the rest wait and share its result.

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

// getUser returns the user data for id, using rdb as a cache and falling
// back to the database on a cache miss. Concurrent calls for the same id
// share a single cache-miss/database round trip.
func getUser(ctx context.Context, rdb *redis.Client, id string) (string, error) {
	key := "user:" + id

	v, err, _ := group.Do(key, func() (interface{}, error) {
		// Try the cache first.
		val, err := rdb.Get(ctx, key).Result()
		if err == nil {
			return val, nil
		}
		if !errors.Is(err, redis.Nil) {
			return "", err
		}

		// Cache miss: fetch from the database.
		val, err = queryUserFromDB(id)
		if err != nil {
			return "", err
		}

		// Populate the cache for subsequent requests.
		if err := rdb.Set(ctx, key, val, 5*time.Minute).Err(); err != nil {
			return "", err
		}
		return val, nil
	})
	if err != nil {
		return "", err
	}

	return v.(string), nil
}

// queryUserFromDB simulates a slow database query.
func queryUserFromDB(id string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return fmt.Sprintf("user-data-for-%s", id), nil
}
```

If a hundred requests for `/users/42` arrive while the cache entry is missing or expired, only one of them queries the database and writes to Redis; the other ninety-nine wait on `group.Do` and receive the same value.

## Comparison with related patterns

- **vs [Mutex]({{ site.baseurl }}/sync/mutex)**.
A mutex serializes *all* calls to a critical section, one at a time, regardless of what they're doing. Singleflight only coalesces calls that share the *same key* — requests for different keys still execute concurrently.

- **vs [`sync.Once`]({{ site.baseurl }}/creational/lazy-initialization)**.
`sync.Once` guarantees a piece of work runs exactly once for the lifetime of the program. Singleflight has no such permanence: once the in-flight call for a key completes, the next call for that key triggers a fresh execution — it deduplicates *concurrent* calls, not *all future* calls.

- **vs [Caching]({{ site.baseurl }}/stability/caching) (TTL cache)**.
A cache avoids *repeated* work across time by storing results. Singleflight avoids *redundant concurrent* work at a single point in time. They solve different problems and are often combined: singleflight protects the cache-population step so that, on a cache miss, only one goroutine recomputes the value while the rest wait and then read from the cache.
