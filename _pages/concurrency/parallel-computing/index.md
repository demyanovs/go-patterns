---
layout: default
title: Parallel Computing
description: "Learn how to use Go’s concurrency primitives to perform efficient parallel computing tasks."
nav_order: 3
permalink: /parallel-computing
parent: Concurrency Patterns
has_children: true
---

# Parallel Computing Patterns

Patterns that involve running multiple independent tasks concurrently to improve performance.

## 1. [Worker Pool]({{ site.baseurl }}/parallel-computing/worker-pool)
Limits concurrency by reusing a fixed number of goroutines to process tasks.

## 2. [Queuing]({{ site.baseurl }}/parallel-computing/queuing)
Buffers tasks for controlled, sequential or concurrent processing.

## 3. [Parallel For Loop]({{ site.baseurl }}/parallel-computing/parallel-for-loop)
Runs loop iterations concurrently using goroutines.

## 4. [Map-Reduce]({{ site.baseurl }}/parallel-computing/map-reduce)
Distributes tasks across multiple workers and aggregates results.

## 5. [Future (Promise)]({{ site.baseurl }}/parallel-computing/future)
Represents a value that will be available at some point in the future, allowing asynchronous computation.

## 6. [Error Group]({{ site.baseurl }}/parallel-computing/errgroup)
Runs goroutines in parallel with error handling and cancellation.