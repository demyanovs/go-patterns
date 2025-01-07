package main

import (
	"log"
	"time"
)

// Functional Options

/*
Now, given this scenario, how do we extend configuration options for our server? There are a few options

Declare new a constructor for each different configuration option
Define a new Config struct that holds configuration information
Use the Functional Option Pattern
*/

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func main() {
	svr := New(
		WithHost("localhost"),
		WithPort(8080),
		WithTimeout(time.Minute),
		WithMaxConn(120),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}

func New(options ...func(*Server)) *Server {
	svr := &Server{}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func (s *Server) Start() error {
	// todo
	return nil
}

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) func(*Server) {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}
