---
layout: default
title: Parallel Computing
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

## 3. [Parallel for loop]({{ site.baseurl }}/parallel-computing/parallel-for-loop)
Runs loop iterations concurrently using goroutines.

## 4. [ErrGroup]({{ site.baseurl }}/parallel-computing/errgroup)
Runs goroutines in parallel with error handling and cancellation.