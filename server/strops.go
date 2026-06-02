package server

import (
	"errors"

	"github.com/nilayrajderkar/redis-implementation/resp"
)

// Append appends suffix to the string value stored at key and returns the
// new length of the value. If the key does not exist, it is created with
// suffix as its value (like Redis APPEND).
func (s *Store) Append(key, suffix string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	existing := s.data[key]
	newValue := existing.value + suffix
	s.data[key] = entry{value: newValue}
	return len(newValue)
}

// handleAppend implements APPEND key value.
func handleAppend(args []interface{}) string {
	if len(args) < 1 {
		return resp.Serialize(errors.New("APPEND requires a key and a value"))
	}

	key, ok := args[0].(*string)
	if !ok {
		return resp.Serialize(errors.New("APPEND key must be a string"))
	}

	value, ok := args[1].(*string)
	if !ok {
		return resp.Serialize(errors.New("APPEND value must be a string"))
	}

	n := DefaultStore.Append(*key, *value)
	return resp.Serialize(n)
}
