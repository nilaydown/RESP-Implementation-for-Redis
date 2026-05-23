package server

import (
	"testing"
	"time"
)

func TestPersistRemovesTTL(t *testing.T) {
	s := NewStore()
	s.Set("k", "v", 200*time.Millisecond)

	if ok := s.Persist("k"); !ok {
		t.Fatalf("expected persist to return true")
	}
	if ttl := s.TTL("k"); ttl != -1 {
		t.Fatalf("expected TTL -1 after persist, got %d", ttl)
	}
}

func TestPersistNoTTL(t *testing.T) {
	s := NewStore()
	s.Set("k", "v", 0)

	if ok := s.Persist("k"); ok {
		t.Fatalf("expected persist to return false for key without TTL")
	}
}

func TestPersistExpiredKey(t *testing.T) {
	s := NewStore()
	s.Set("k", "v", 50*time.Millisecond)
	time.Sleep(100 * time.Millisecond)

	if ok := s.Persist("k"); ok {
		t.Fatalf("expected persist to return false for expired key")
	}
	if ttl := s.TTL("k"); ttl != -2 {
		t.Fatalf("expected TTL -2 for missing/expired key, got %d", ttl)
	}
}
