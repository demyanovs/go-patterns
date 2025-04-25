---
layout: home
title: Concurrency Patterns
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
  - [Error Group]({{ site.baseurl }}/parallel-computing/errgroup)
- [Delayed Computing]({{ site.baseurl }}/delayed-computing)
  - [Future (Promise)]({{ site.baseurl }}/delayed-computing/future)

![Concurrency Patterns]({{ site.baseurl }}/assets/images/concurrent_patterns_01.png)