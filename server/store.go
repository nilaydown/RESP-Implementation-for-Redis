package server

import (
	"fmt"
	"strconv"
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

// Keys returns all keys matching a glob-like pattern. Supports * as wildcard.
// An empty pattern or "*" returns all non-expired keys.
func (s *Store) Keys(pattern string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	now := time.Now()
	var result []string
	for k, e := range s.data {
		if e.hasTTL && now.After(e.expiresAt) {
			continue
		}
		if pattern == "" || pattern == "*" || matchGlob(pattern, k) {
			result = append(result, k)
		}
	}
	return result
}

func matchGlob(pattern, s string) bool {
	pi, si := 0, 0
	starPi, starSi := -1, -1
	for si < len(s) {
		if pi < len(pattern) && (pattern[pi] == s[si] || pattern[pi] == '?') {
			pi++
			si++
		} else if pi < len(pattern) && pattern[pi] == '*' {
			starPi = pi
			starSi = si
			pi++
		} else if starPi >= 0 {
			pi = starPi + 1
			starSi++
			si = starSi
		} else {
			return false
		}
	}
	for pi < len(pattern) && pattern[pi] == '*' {
		pi++
	}
	return pi == len(pattern)
}

// Incr atomically increments the integer value at key by delta.
// If the key does not exist, it is initialized to 0 before the operation.
// Returns the new value or an error if the current value is not an integer.
func (s *Store) Incr(key string, delta int) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	current := 0
	if e, ok := s.data[key]; ok {
		if e.hasTTL && time.Now().After(e.expiresAt) {
			delete(s.data, key)
		} else {
			val, err := strconv.Atoi(e.value)
			if err != nil {
				return 0, fmt.Errorf("value is not an integer or out of range")
			}
			current = val
		}
	}

	current += delta
	s.data[key] = entry{value: strconv.Itoa(current)}
	return current, nil
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
