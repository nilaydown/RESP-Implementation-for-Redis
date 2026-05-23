package server

import (
	"testing"
	"time"
)

func TestPersistRemovesTTL(t *testing.T) {
	s := NewStore()
	s.Set("tmp", "val", 1*time.Second)

	if ok := s.Persist("tmp"); !ok {
		t.Fatalf("expected Persist to return true for key with TTL")
	}
	if ttl := s.TTL("tmp"); ttl != -1 {
		t.Fatalf("expected TTL -1 after persist, got %d", ttl)
	}
}

func TestPersistNoTTL(t *testing.T) {
	s := NewStore()
	s.Set("perm", "val", 0)

	if ok := s.Persist("perm"); ok {
		t.Fatalf("expected Persist to return false for key without TTL")
	}
}

func TestPersistExpiredKey(t *testing.T) {
	s := NewStore()
	s.Set("exp", "val", 50*time.Millisecond)
	time.Sleep(100 * time.Millisecond)

	if ok := s.Persist("exp"); ok {
		t.Fatalf("expected Persist to return false for expired/missing key")
	}
	if ttl := s.TTL("exp"); ttl != -2 {
		t.Fatalf("expected TTL -2 for missing key, got %d", ttl)
	}
}
