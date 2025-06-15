---
layout: home
title: Concurrency Patterns
description: "Explore foundational concurrency patterns in Go, including channel pipelines, fan-in, fan-out, and parallel processing."
nav_order: 2
permalink: /concurrency
has_children: true
---

# Concurrency Patterns
Patterns address managing multiple tasks concurrently, leveraging its built-in goroutines and channels for efficient parallelism.

Explore categories:
- [Generative]({{ site.baseurl }}/generative)
  - [Generator]({{ site.baseurl }}/generative/generator)
  - [Fan In]({{ site.baseurl }}/generative/fan-in)
  - [Fan Out]({{ site.baseurl }}/generative/fan-out)
  - [Pipeline]({{ site.baseurl }}/generative/pipeline)
- [Synchronization]({{ site.baseurl }}/sync)
  - [Mutex]({{ site.baseurl }}/sync/mutex)
  - [Semaphore]({{ site.baseurl }}/sync/semaphore)
- [Parallel Computing]({{ site.baseurl }}/parallel-computing)
  - [Worker Pool]({{ site.baseurl }}/parallel-computing/worker-pool)
  - [Queuing]({{ site.baseurl }}/parallel-computing/queuing)
  - [Parallel For Loop]({{ site.baseurl }}/parallel-computing/parallel-for-loop)
  - [Map-Reduce]({{ site.baseurl }}/parallel-computing/map-reduce)
  - [Future (Promise)]({{ site.baseurl }}/parallel-computing/future)
  - [Error Group]({{ site.baseurl }}/parallel-computing/errgroup)

![Concurrency Patterns]({{ site.baseurl }}/assets/images/concurrent_patterns_01.png)