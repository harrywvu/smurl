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

func (s *Store) Create(code, url string) bool {

	s.mu.Lock()
	defer s.mu.Unlock()

	// check if code already exists = false
	if _, exists := s.shortURLs[code]; exists {
		return false
	}

	s.shortURLs[code] = url

	return true

}

func (s *Store) Resolve(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	// lookup map, if exists, return url and true
	// if _, exists := s.shortURLs[code]; exists {
	// 	return s.shortURLs[code], true
	// }

	// return "", false

	url, exists := s.shortURLs[code]
	return url, exists
}