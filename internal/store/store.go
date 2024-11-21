package store

import (
	"errors"
	"sync"
)

type KVStore interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Persist() error
	Load() error
}

type InMemoryKVStore struct {
	store map[string][]byte
	mu    sync.RWMutex
}

func NewInMemoryKVStore() *InMemoryKVStore {
	return &InMemoryKVStore{
		store: make(map[string][]byte),
	}
}

func (s *InMemoryKVStore) Set(key string, value []byte) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[key] = value
	return nil
}

func (s *InMemoryKVStore) Get(key string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exists := s.store[key]
	if !exists {
		return nil, errors.New("key not found")
	}

	return value, nil
}

func (s *InMemoryKVStore) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.store[key]; !exists {
		return errors.New("key not found")
	}
	delete(s.store, key)
	return nil
}
