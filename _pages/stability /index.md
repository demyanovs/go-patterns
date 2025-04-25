---
layout: default
title: Stability Patterns
nav_order: 4
permalink: /stability

---

# Stability Patterns
Patterns focus on making systems more resilient, ensuring they handle failure gracefully.

## 1. Circuit Breaker
Prevents repeated calls to a failing service, avoiding cascading failures.

## 2. Retry
Retries failed operations with optional backoff strategies.

## 3. Timeout
Prevents operations from running indefinitely by enforcing time limits.

## 4. Rate Limiting
Controls the rate of requests to prevent system overload.

## 5. Caching
Reduce repeated expensive computations or I/O.

![Concurrency Patterns]({{ site.baseurl }}/assets/images/stability_patterns_01.png)