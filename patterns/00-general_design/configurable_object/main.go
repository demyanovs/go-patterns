package main

import (
	"flag"
	"os"
)

// Builder Pattern
// Configurable Object, методы конфигураторы которого чейнятся. Получится что-то вроде.

type Config struct {
	Host  string
	Port  int
	Debug bool
	// ...
}

type ConfigBuilder struct {
	config Config
}

func (b ConfigBuilder) WithHost(host string) ConfigBuilder {
	b.config.Host = host
	return b
}

func (b ConfigBuilder) WithPort(port int) ConfigBuilder {
	b.config.Port = port
	return b
}

func (b ConfigBuilder) WithDebug(debug bool) ConfigBuilder {
	b.config.Debug = debug
	return b
}

func NewConfigFromFlags() (Config, error) {
	var port int
	flag.IntVar(&port, "port", 8080, "server port")

	var host string
	flag.StringVar(&host, "host", "", "server host")

	var debug bool
	flag.BoolVar(&debug, "debug", false, "enable debugging")

	flag.Parse()

	// create builder
	var builder ConfigBuilder

	// customize configuration using flags
	builder.
		WithHost(host).
		WithPort(port).
		WithDebug(debug)

	// customize configuration using environment variables
	if envHost := os.Getenv("HOST"); envHost != "" {
		builder.WithHost(envHost)
	}

	return builder.Config, nil
}
