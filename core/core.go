package core

import (
	"errors"
	"sync"
)

type Storage struct {
	sync.RWMutex
	m map[string]string
}

var ErrNotFound = errors.New("no key found")

func NewStorage() *Storage {
	return &Storage{m: make(map[string]string)}
}

func (s *Storage) Put(key string, value string) error {
	s.RWMutex.Lock()
	s.m[key] = value
	s.RWMutex.Unlock()
	return nil
}

func (s *Storage) Get(key string) (string, error) {
	s.RWMutex.RLock()
	value, ok := s.m[key]
	s.RWMutex.RUnlock()
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (s *Storage) Delete(key string) error {
	delete(s.m, key)
	return nil
}
