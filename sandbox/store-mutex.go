package main

import (
	"errors"
	"sync"
)

var ErrorNoSuchKey = errors.New("no such key")

type store struct {
	sync.RWMutex
	m map[string]string
}

func (s *store) get(key string) (string, error) {
	s.RWMutex.RLocker()
	defer s.RWMutex.RUnlock()

	val, ok := s.m[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return val, nil
}

func (s *store) put(key string, value string) error {
	s.Lock()
	defer s.Unlock()

	s.m[key] = value

	return nil
}
