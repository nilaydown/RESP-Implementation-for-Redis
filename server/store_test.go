//go:build ignore

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
}

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

func TestGetMissing(t *testing.T) {
	s := NewStore()
	_, ok := s.Get("nope")
	if ok {
		t.Fatal("expected key to not exist")
	}

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
}

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

func TestOverwrite(t *testing.T) {
	s := NewStore()
	s.Set("k", "v1", 0)
	s.Set("k", "v2", 0)

	val, _ := s.Get("k")
	if val != "v2" {
		t.Fatalf("expected v2, got %s", val)
	}

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
}

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

func TestDel(t *testing.T) {
	s := NewStore()
	s.Set("a", "1", 0)
	s.Set("b", "2", 0)

	deleted := s.Del("a", "b", "c")
	if deleted != 2 {
		t.Fatalf("expected 2 deleted, got %d", deleted)
	}

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
	if _, ok := s.Get("a"); ok {
		t.Fatal("a should be deleted")
	}

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
}

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

func TestExists(t *testing.T) {
	s := NewStore()
	s.Set("x", "1", 0)

	if s.Exists("x") != 1 {
		t.Fatal("x should exist")
	}

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
	if s.Exists("y") != 0 {
		t.Fatal("y should not exist")
	}

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
	if s.Exists("x", "y") != 1 {
		t.Fatal("only x should exist")
	}

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
}

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

func TestTTLExpiry(t *testing.T) {
	s := NewStore()
	s.Set("temp", "val", 100*time.Millisecond)

	if _, ok := s.Get("temp"); !ok {
		t.Fatal("key should exist before expiry")
	}

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

	time.Sleep(150 * time.Millisecond)

	if _, ok := s.Get("temp"); ok {
		t.Fatal("key should have expired")
	}

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
}

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

func TestIncrNewKey(t *testing.T) {
	s := NewStore()
	val, err := s.Incr("counter", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

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
	if val != 1 {
		t.Fatalf("expected 1, got %d", val)
	}

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
}

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

func TestIncrExistingKey(t *testing.T) {
	s := NewStore()
	s.Set("counter", "10", 0)

	val, err := s.Incr("counter", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

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
	if val != 11 {
		t.Fatalf("expected 11, got %d", val)
	}

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
}

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

func TestDecrKey(t *testing.T) {
	s := NewStore()
	s.Set("counter", "10", 0)

	val, err := s.Incr("counter", -1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

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
	if val != 9 {
		t.Fatalf("expected 9, got %d", val)
	}

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
}

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

func TestIncrNonInteger(t *testing.T) {
	s := NewStore()
	s.Set("name", "alice", 0)

	_, err := s.Incr("name", 1)
	if err == nil {
		t.Fatal("expected error for non-integer value")
	}

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
}

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

func TestIncrExpiredKey(t *testing.T) {
	s := NewStore()
	s.Set("counter", "10", 50*time.Millisecond)

	time.Sleep(100 * time.Millisecond)

	val, err := s.Incr("counter", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

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
	if val != 1 {
		t.Fatalf("expected 1 (expired key treated as new), got %d", val)
	}

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
}

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

func TestTTLValues(t *testing.T) {
	s := NewStore()

	if s.TTL("missing") != -2 {
		t.Fatal("missing key should return -2")
	}

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

	s.Set("perm", "val", 0)
	if s.TTL("perm") != -1 {
		t.Fatal("key without TTL should return -1")
	}

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

	s.Set("tmp", "val", 10*time.Second)
	ttl := s.TTL("tmp")
	if ttl < 8 || ttl > 10 {
		t.Fatalf("expected TTL ~10, got %d", ttl)
	}

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
}

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

func TestKeysAll(t *testing.T) {
	s := NewStore()
	s.Set("a", "1", 0)
	s.Set("b", "2", 0)
	s.Set("c", "3", 0)

	keys := s.Keys("*")
	if len(keys) != 3 {
		t.Fatalf("expected 3 keys, got %d", len(keys))
	}

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
}

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

func TestKeysPattern(t *testing.T) {
	s := NewStore()
	s.Set("user:1", "alice", 0)
	s.Set("user:2", "bob", 0)
	s.Set("session:1", "data", 0)

	keys := s.Keys("user:*")
	if len(keys) != 2 {
		t.Fatalf("expected 2 keys matching user:*, got %d", len(keys))
	}

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
}

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

func TestKeysExcludesExpired(t *testing.T) {
	s := NewStore()
	s.Set("alive", "yes", 0)
	s.Set("dying", "soon", 50*time.Millisecond)

	time.Sleep(100 * time.Millisecond)

	keys := s.Keys("*")
	if len(keys) != 1 {
		t.Fatalf("expected 1 key (expired excluded), got %d", len(keys))
	}

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
}

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
