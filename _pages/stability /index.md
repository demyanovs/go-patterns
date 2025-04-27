---
layout: default
title: Stability Patterns
description: "Implement patterns like circuit breakers and retries to build resilient and fault-tolerant Go systems."
nav_order: 4
permalink: /stability
has_children: true
---

# Stability Patterns
Patterns focus on making systems more resilient, ensuring they handle failure gracefully.

## 1. Circuit Breaker
Prevents repeated calls to a failing service, avoiding cascading failures.

## 2. [Retry]({{ site.baseurl }}/stability/retry)
Retries failed operations with optional backoff strategies.

## 3. [Timeout]({{ site.baseurl }}/stability/timeout)
Prevents operations from running indefinitely by enforcing time limits.

## 5. [Caching]({{ site.baseurl }}/stability/caching)
Reduces repeated expensive computations or I/O.
 - [With Map and sync.RWMutex]({{ site.baseurl }}/stability/caching/with-map-and-rwmutex)
 - [With sync.Map]({{ site.baseurl }}/stability/caching/with-sync-map)
 - [With Automatic Cleanup]({{ site.baseurl }}/stability/caching/with-expiration)

![Concurrency Patterns]({{ site.baseurl }}/assets/images/stability_patterns_01.png)