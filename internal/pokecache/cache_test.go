package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "pikachu"
	val := []byte("electric pokemon")

	cache.Add(key, val)

	got, found := cache.Get(key)
	if !found {
		t.Errorf("Expected key to be found")
	}

	if !bytes.Equal(got, val) {
		t.Errorf("expected %s, got %s", string(val), string(got))
	}
}

func TestGetMissing(t *testing.T) {
	cache := NewCache(5 * time.Second)

	got, found := cache.Get("charizard")
	if found {
		t.Errorf("Expected key to not be found")
	}

	if got != nil {
		t.Errorf("Expected key to be nil")
	}
}
