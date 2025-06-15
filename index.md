---
layout: default
title: Home
nav_order: 1
---

# Welcome to Go Patterns

Welcome to the collection of idiomatic Go design patterns! This site gathers a variety of Go patterns, each with clear, practical code examples.

Designed for clarity, learning, and collaboration.
Inspired by effective Go idioms and concurrency best practices.

Explore categories:
- [Concurrency Patterns]({{ site.baseurl }}/concurrency)
  - [Generative]({{ site.baseurl }}/generative)
    - [Generator]({{ site.baseurl }}/generative/generator)
    - [Fan In]({{ site.baseurl }}/generative/fan-in)
    - [Fan Out]({{ site.baseurl }}/generative/fan-out)
    - [Pipeline]({{ site.baseurl }}/generative/pipeline)
  - [Synchronization Patterns]({{ site.baseurl }}/sync)
    - [Mutex (Rate Limiting)]({{ site.baseurl }}/sync/mutex)
    - [Semaphore (Rate Limiting)]({{ site.baseurl }}/sync/semaphore)
  - [Parallel Computing Patterns]({{ site.baseurl }}/parallel-computing)
    - [Worker Pool]({{ site.baseurl }}/parallel-computing/worker-pool)
    - [Queuing]({{ site.baseurl }}/parallel-computing/queuing)
    - [Parallel For Loop]({{ site.baseurl }}/parallel-computing/parallel-for-loop)
    - [Map-Reduce]({{ site.baseurl }}/parallel-computing/map-reduce)
    - [Future (Promise)]({{ site.baseurl }}/parallel-computing/future)
    - [Error Group]({{ site.baseurl }}/parallel-computing/errgroup)
- [Creational Patterns]({{ site.baseurl }}/creational)
  - [Functional Options]({{ site.baseurl }}/creational/functional-options)
  - [Configuration Struct + Factory]({{ site.baseurl }}/creational/configuration-struct-factory)
  - [Configurable Object]({{ site.baseurl }}/creational/configurable-object)
  - [Lazy Initialization]({{ site.baseurl }}/creational/lazy-initialization)
  - [Fluent Interfaces]({{ site.baseurl }}/creational/fluent-interfaces)
- [Stability Patterns]({{ site.baseurl }}/stability)
  - [Retry]({{ site.baseurl }}/stability/retry)
  - [Timeout]({{ site.baseurl }}/stability/timeout)
  - [Drop]({{ site.baseurl }}/stability/drop)
  - [Circuit Breaker]({{ site.baseurl }}/stability/circuit-breaker)
  - [Caching]({{ site.baseurl }}/stability/caching)
    - [With Map and sync.RWMutex]({{ site.baseurl }}/stability/caching/with-map-and-rwmutex)
    - [With sync.Map]({{ site.baseurl }}/stability/caching/with-sync-map)
    - [With Automatic Cleanup]({{ site.baseurl }}/stability/caching/with-expiration)

## Articles
  - [Channels]({{ site.baseurl }}/articles/channels)