package main

import "sync"

type Store struct{
	mu sync.RWMutex
	shortURLs map[string]string
}

// initialize the map
func NewStore() *Store {
	return &Store{shortURLs: make(map[string]string)}
}

// create a key-value pair in the store
func (s *Store) Create(code, url string) bool {

	s.mu.Lock()
	defer s.mu.Unlock()

	// check if code already exists = false
	if _, exists := s.shortURLs[code]; exists {
		return false
	}

	// add code in store
	s.shortURLs[code] = url

	return true

}

// returns the original url from the code = true
func (s *Store) Resolve(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, exists := s.shortURLs[code]
	return url, exists
}