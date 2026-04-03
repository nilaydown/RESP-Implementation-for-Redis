package server

import (
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {
	s := NewStore()
	s.Set("name", "alice", 0)

	val, ok := s.Get("name")
	if !ok || val != "alice" {
		t.Fatalf("expected alice, got %q (ok=%v)", val, ok)
	}
}

func TestGetMissing(t *testing.T) {
	s := NewStore()
	_, ok := s.Get("nope")
	if ok {
		t.Fatal("expected key to not exist")
	}
}

func TestOverwrite(t *testing.T) {
	s := NewStore()
	s.Set("k", "v1", 0)
	s.Set("k", "v2", 0)

	val, _ := s.Get("k")
	if val != "v2" {
		t.Fatalf("expected v2, got %s", val)
	}
}

func TestDel(t *testing.T) {
	s := NewStore()
	s.Set("a", "1", 0)
	s.Set("b", "2", 0)

	deleted := s.Del("a", "b", "c")
	if deleted != 2 {
		t.Fatalf("expected 2 deleted, got %d", deleted)
	}
	if _, ok := s.Get("a"); ok {
		t.Fatal("a should be deleted")
	}
}

func TestExists(t *testing.T) {
	s := NewStore()
	s.Set("x", "1", 0)

	if s.Exists("x") != 1 {
		t.Fatal("x should exist")
	}
	if s.Exists("y") != 0 {
		t.Fatal("y should not exist")
	}
	if s.Exists("x", "y") != 1 {
		t.Fatal("only x should exist")
	}
}

func TestTTLExpiry(t *testing.T) {
	s := NewStore()
	s.Set("temp", "val", 100*time.Millisecond)

	if _, ok := s.Get("temp"); !ok {
		t.Fatal("key should exist before expiry")
	}

	time.Sleep(150 * time.Millisecond)

	if _, ok := s.Get("temp"); ok {
		t.Fatal("key should have expired")
	}
}

func TestTTLValues(t *testing.T) {
	s := NewStore()

	if s.TTL("missing") != -2 {
		t.Fatal("missing key should return -2")
	}

	s.Set("perm", "val", 0)
	if s.TTL("perm") != -1 {
		t.Fatal("key without TTL should return -1")
	}

	s.Set("tmp", "val", 10*time.Second)
	ttl := s.TTL("tmp")
	if ttl < 8 || ttl > 10 {
		t.Fatalf("expected TTL ~10, got %d", ttl)
	}
}
