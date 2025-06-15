---
layout: default
title: Caching Patterns
description: "Optimize performance in Go applications with caching strategies for repeated data access."
nav_order: 6
parent: Stability Patterns
permalink: /stability/caching
has_children: true
---

# Caching Patterns

Caching Patterns store expensive or frequently requested data in memory to improve performance and reduce repeated computations or data fetching.

## 1. [With Map and sync.RWMutex]({{ site.baseurl }}/stability/caching/with-map-and-rwmutex)
Thread-safe cache with manual locking.

## 2. [With sync.Map]({{ site.baseurl }}/stability/caching/with-sync-map)
Concurrent cache with built-in locking.

## 3. [With Automatic Cleanup]({{ site.baseurl }}/stability/caching/with-expiration)
Cache with background removal of stale data.
