package server

import (
	"errors"

	"github.com/nilayrajderkar/redis-implementation/resp"
)

// GetSet atomically sets key to newValue and returns the value that was
// previously stored there (and whether a previous value existed), matching
// Redis GETSET.
func (s *Store) GetSet(key, newValue string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	previous, existed := s.data[key]
	s.data[key] = entry{value: newValue}
	return previous.value, existed
}

// handleGetSet implements GETSET key value.
func handleGetSet(args []interface{}) string {
	if len(args) < 1 {
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
		return resp.Serialize(nil)
	}
	return resp.Serialize(old)
}
