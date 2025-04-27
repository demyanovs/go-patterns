---
layout: default
title: Generative Patterns
description: "Learn how generative concurrency patterns create and manage Go routines and channels efficiently."
nav_order: 1
permalink: /generative
parent: Concurrency Patterns
has_children: true
---

# Generative Patterns

Patterns focused on spawning new goroutines, often involving channels or dynamic generation.

## 1. [Generator]({{ site.baseurl }}/generative/generator)
Simple value producer using goroutines and channels.

## 2. [Fan In]({{ site.baseurl }}/generative/fan-in)
Merging multiple inputs into one output.

## 3. [Fan Out]({{ site.baseurl }}/generative/fan-out)
Splitting one input to multiple outputs.

## 4. [Pipeline]({{ site.baseurl }}/generative/pipeline)
Series of stages connected by channels.
