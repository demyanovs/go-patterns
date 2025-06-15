---
layout: default
title: Lazy Initialization
description: "Lazy Initialization pattern in Go: Delaying object creation until it's needed."
nav_order: 4
parent: Creational Patterns
permalink: /creational/lazy-initialization
---

# Lazy Initialization

**Lazy Initialization** is a pattern where an object or resource is not created until it is actually needed. 
This helps improve performance and resource usage, especially when the creation is expensive and not always required.

## Applicability

 - Creating an object is expensive and may not be needed. 
 - Delay resource allocation until first use. 
 - Thread-safe, one-time initialization.

## Example

```go
package main

import (
	"fmt"
	"sync"
)

type Config struct {
	data string
}

var (
	config     *Config
	configOnce sync.Once
)

func GetConfig() *Config {
	configOnce.Do(func() {
		fmt.Println("Initializing config...")
		config = &Config{data: "Loaded configuration"}
	})
	return config
}

func main() {
	fmt.Println("First call:")
	cfg1 := GetConfig()
	fmt.Println(cfg1.data)

	fmt.Println("\nSecond call:")
	cfg2 := GetConfig()
	fmt.Println(cfg2.data)
}
```

### Output
```
First call:
Initializing config...
Loaded configuration

Second call:
Loaded configuration
```
