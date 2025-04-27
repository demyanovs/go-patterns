---
layout: default
title: Functional Options
description: "Build highly configurable Go structs and constructors using the Functional Options pattern."
nav_order: 1
parent: Creational Patterns
permalink: /creational/functional-options
---

# Functional Options

The **Functional Options** pattern is flexible and extensible approach, where functions (often called "option functions") are used to configure an object. 
Each option function takes a pointer to the object and modifies it, allowing for a chainable and fluent interface. 
This is often preferred when you have multiple optional parameters and want to avoid a large constructor function with many parameters.

In Go, this is typically implemented by passing functions that set fields of a struct, rather than using setter methods.

## Applicability

 - **There are many optional parameters**.<br/>
Instead of overloading constructors or using large config structs, use option functions for clarity.

 - **To enforce immutability after creation**.<br/>
The object can be created in a fully configured state with no need for mutable setters afterward.

- **The configuration logic is complex or reusable**.<br/>
Each option function can encapsulate reusable logic, e.g., validation or derived values.

- **Provide a clean and fluent API**.<br/>
Makes the construction of objects readable, composable, and extendable.

- **Backward compatibility is important**.<br/>
Adding a new option doesn’t change the constructor’s signature, reducing the chance of breaking changes.

## Example

```go
package main

type Database struct {
	host     string
	port     int
	username string
	password string
}

type Option func(*Database)

func NewDatabase(options ...Option) *Database {
	db := &Database{}
	for _, option := range options {
		option(db)
	}
	return db
}

func WithHost(host string) Option {
	return func(db *Database) {
		db.host = host
	}
}

func WithPort(port int) Option {
	return func(db *Database) {
		db.port = port
	}
}

func WithUsername(username string) Option {
	return func(db *Database) {
		db.username = username
	}
}

func WithPassword(password string) Option {
	return func(db *Database) {
		db.password = password
	}
}

```