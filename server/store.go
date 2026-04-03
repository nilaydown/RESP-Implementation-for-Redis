package server

import (
	"sync"
	"time"
)

type entry struct {
	value     string
	expiresAt time.Time
	hasTTL    bool
}

// Store is a thread-safe in-memory key-value store with TTL support.
type Store struct {
	mu   sync.RWMutex
	data map[string]entry
}

// NewStore creates an empty Store and starts a background cleanup goroutine
// that evicts expired keys every second.
func NewStore() *Store {
	s := &Store{data: make(map[string]entry)}
	go s.evictLoop()
	return s
}

func (s *Store) Set(key, value string, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	e := entry{value: value}
	if ttl > 0 {
		e.expiresAt = time.Now().Add(ttl)
		e.hasTTL = true
	}
	s.data[key] = e
}

// Get returns the value and true if the key exists and hasn't expired,
// otherwise ("", false).
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	e, ok := s.data[key]
	s.mu.RUnlock()

	if !ok {
		return "", false
	}
	if e.hasTTL && time.Now().After(e.expiresAt) {
		s.Del(key)
		return "", false
	}
	return e.value, true
}

// Del removes the given keys and returns how many were actually deleted.
func (s *Store) Del(keys ...string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	deleted := 0
	for _, k := range keys {
		if _, ok := s.data[k]; ok {
			delete(s.data, k)
			deleted++
		}
	}
	return deleted
}

// Exists returns the number of given keys that exist (and haven't expired).
func (s *Store) Exists(keys ...string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	count := 0
	now := time.Now()
	for _, k := range keys {
		if e, ok := s.data[k]; ok {
			if !e.hasTTL || now.Before(e.expiresAt) {
				count++
			}
		}
	}
	return count
}

// TTL returns the remaining TTL in seconds. -2 means key doesn't exist,
// -1 means no expiry set.
func (s *Store) TTL(key string) int {
	s.mu.RLock()
	e, ok := s.data[key]
	s.mu.RUnlock()

	if !ok {
		return -2
	}
	if !e.hasTTL {
		return -1
	}
	remaining := time.Until(e.expiresAt)
	if remaining <= 0 {
		s.Del(key)
		return -2
	}
	return int(remaining.Seconds())
}

func (s *Store) evictLoop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		s.mu.Lock()
		now := time.Now()
		for k, e := range s.data {
			if e.hasTTL && now.After(e.expiresAt) {
				delete(s.data, k)
			}
		}
		s.mu.Unlock()
	}
}
