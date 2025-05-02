---
layout: default
title: Creational Patterns
description: "Discover Go design patterns for creating flexible, reusable, and testable object construction logic."
nav_order: 3
permalink: /creational
has_children: true
---

# Creational Patterns
Patterns focus on flexible and controlled object creation, often using interfaces, functions, and struct composition rather than traditional class-based inheritance.

## 1. [Functional Options]({{ site.baseurl }}/creational/functional-options)
Uses variadic option functions to configure an object.

## 2. [Configuration Struct + Factory]({{ site.baseurl }}/creational/configuration-struct-factory)
Mimics traditional builders with method chaining.

## 3. [Configurable Object]({{ site.baseurl }}/creational/configurable-object)
Allows an object to be created with default values and then configured step-by-step through setter methods.

## 4. [Lazy Initialization]({{ site.baseurl }}/creational/lazy-initialization)
Delays object creation until it's needed.

## 5. [Fluent Interfaces]({{ site.baseurl }}/creational/fluent-interfaces)
Chaining method calls on the same object, often seen in SQL builders or HTTP clients.

![Creational Patterns]({{ site.baseurl }}/assets/images/creational_patterns_01.png)

