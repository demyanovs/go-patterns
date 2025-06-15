---
layout: default
title: Configuration Struct + Factory
description: "Create scalable object factories in Go using configuration structs for flexibility and clarity."
nav_order: 2
parent: Creational Patterns
permalink: /creational/configuration-struct-factory
---

# Configuration Struct + Factory

The **Configuration Struct + Factory** pattern uses a struct to hold configuration options and a factory function to create and initialize the object.

## Applicability

 - When configuration values are grouped and passed together.
 - When you want **simple and readable** object creation.
 - When you need to **validate or process** config before object creation.
 - Useful for passing options across layers or packages.

## Example

```go
package main

type Config struct {
	Host string
	Port int
}

type Server struct {
	host string
	port int
}

func NewServer(cfg Config) *Server {
	return &Server{
		host: cfg.Host,
		port: cfg.Port,
	}
}
```