---
layout: default
title: Configurable Object
description: "Implement configurable objects in Go to allow flexible runtime behavior without complex constructors."
nav_order: 4
parent: Creational Patterns
permalink: /creational/configurable-object
---

# Configurable Object

The **Configurable Object** pattern involves creating a struct (or object) with fields that can be modified after initialization. 
These fields are often set through methods that modify the internal state of the object. 
This pattern allows for flexible configuration of an object, where the user can set various parameters in any order.

For example, you might have an object where various options can be set through methods, and those options can be added or changed over time.

## Applicability

 - **Default construction is simple, but configuration is optional**.<br/>
Use when an object can be created with defaults and configured step-by-step using setters.

 - **Clear and readable configuration phase**.<br/>
Especially useful when the object will be configured at different points in time, possibly conditionally.

- **Mutability is acceptable**.<br/>
This pattern assumes the object’s internal state can be modified after creation.

- **To avoid complex constructors**.<br/>
Instead of passing many parameters to a constructor, you provide setters to configure as needed.

## Example

```go
package main

type Database struct {
	host     string
	port     int
	username string
	password string
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) SetHost(host string) {
	db.host = host
}

func (db *Database) SetPort(port int) {
	db.port = port
}

func (db *Database) SetUsername(username string) {
	db.username = username
}

func (db *Database) SetPassword(password string) {
	db.password = password
}
```