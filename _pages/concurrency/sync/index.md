---
layout: default
title: Synchronization
description: "Explore synchronization strategies like mutexes and semaphores to control concurrency in Go programs."
nav_order: 2
permalink: /sync
parent: Concurrency Patterns
has_children: true
---

# Synchronization Patterns

Patterns involving coordinating shared state between goroutines.

## 1. [Mutex]({{ site.baseurl }}/sync/mutex)
Ensures exclusive access to a resource by one goroutine at a time.

## 2. [Semaphore]({{ site.baseurl }}/sync/semaphore)
Controls the number of goroutines allowed to access a resource at once.