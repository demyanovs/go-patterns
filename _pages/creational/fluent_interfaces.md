---
layout: default
title: Fluent Interfaces
description: "Chain method calls gracefully using Fluent Interfaces to simplify Go code APIs."
nav_order: 5
parent: Creational Patterns
permalink: /creational/fluent-interfaces
---

# Fluent Interfaces

A **Fluent Interface** pattern enables method chaining by returning the object itself from methods. 
This results in readable and expressive code, resembling natural language. 
In Go, it’s often used in builders, configurations, or setup code.

## Applicability

 - **Builder patterns**.<br/>
Creating complex objects step-by-step (e.g., SQL queries, HTTP requests).

- **Configuration APIs**.<br/>
Setting up structs with multiple optional parameters.

- **Test DSLs.**<br/>
Creating expressive test cases.

## Example

```go
package main

import "fmt"

type QueryBuilder struct {
	query string
}

func NewQuery() *QueryBuilder {
	return &QueryBuilder{}
}

func (q *QueryBuilder) Select(fields string) *QueryBuilder {
	q.query += "SELECT " + fields + " "
	return q
}

func (q *QueryBuilder) From(table string) *QueryBuilder {
	q.query += "FROM " + table + " "
	return q
}

func (q *QueryBuilder) Where(condition string) *QueryBuilder {
	q.query += "WHERE " + condition + " "
	return q
}

func (q *QueryBuilder) Build() string {
	return q.query
}

func main() {
	q := NewQuery().Select("*").From("users").Where("age > 18").Build()
	fmt.Println(q)
	// Output: SELECT * FROM users WHERE age > 18
}
```