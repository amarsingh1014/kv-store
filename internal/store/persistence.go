package store

import (
	"os"
	"encoding/json"
)

func (s *InMemoryKVStore) Persist() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	file, err := os.Create("store.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(s.store)
}

func (s *InMemoryKVStore) Load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open("store.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File does not exist, nothing to load
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&s.store)
}
