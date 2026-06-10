package server

import (
	"errors"
	"time"

	"github.com/nilayrajderkar/redis-implementation/resp"
)

// GetSet atomically sets key to newValue and returns the value that was
// previously stored there (and whether a previous value existed), matching
// Redis GETSET.
func (s *Store) GetSet(key, newValue string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	prev := ""
	existed := false
	if e, ok := s.data[key]; ok {
		if e.hasTTL && time.Now().After(e.expiresAt) {
			delete(s.data, key)
		} else {
			prev = e.value
			existed = true
		}
	}
	s.data[key] = entry{value: newValue}
	return prev, existed
}

// handleGetSet implements GETSET key value.
func handleGetSet(args []interface{}) string {
	if len(args) < 2 {
		return resp.Serialize(errors.New("GETSET requires a key and a value"))
	}

	key, ok := args[0].(*string)
	if !ok {
		return resp.Serialize(errors.New("GETSET key must be a string"))
	}

	value, ok := args[1].(*string)
	if !ok {
		return resp.Serialize(errors.New("GETSET value must be a string"))
	}

	old, existed := DefaultStore.GetSet(*key, *value)
	if !existed {
		return "$-1\r\n"
	}
	return resp.Serialize(old)
}
