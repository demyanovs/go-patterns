---
layout: default
title: Database Patterns
description: "Database-related design patterns in Go: transactions, units of work, and keeping data changes consistent."
nav_order: 5
permalink: /database
has_children: true
---

# Database Patterns

These patterns help you group database work into **atomic, consistent steps**.

## 1. [Unit of Work (Transaction Manager)]({{ site.baseurl }}/database/unit-of-work)
Groups related reads and writes into a single transactional unit.

![Database Patterns]({{ site.baseurl }}/assets/images/database_patterns_01.png)