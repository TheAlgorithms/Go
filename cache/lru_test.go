package cache_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/cache"
)

func TestLRU(t *testing.T) {
	cache := cache.NewLRU(3)

	t.Run("Test 1: Put number and Get is correct", func(t *testing.T) {
		key, value := "1", 1
		cache.Put(key, value)
		got := cache.Get(key)

		if got != value {
			t.Errorf("expected: %v, got: %v", value, got)
		}
	})

	t.Run("Test 2: Get data not stored in cache should return nil", func(t *testing.T) {
		got := cache.Get("2")
		if got != nil {
			t.Errorf("expected: nil, got: %v", got)
		}
	})

	t.Run("Test 3: Put data over capacity and Get should return nil", func(t *testing.T) {
		cache.Put("2", 2)
		cache.Put("3", 3)
		cache.Put("4", 4)

		got := cache.Get("1")
		if got != nil {
			t.Errorf("expected: nil, got: %v", got)
		}
	})

	t.Run("test 4: Put key over capacity but recent key exists", func(t *testing.T) {
		cache.Put("4", 4)
		cache.Put("3", 3)
		cache.Put("2", 2)
		cache.Put("1", 1)

		got := cache.Get("4")
		if got != nil {
			t.Errorf("expected: nil, got: %v", got)
		}

		expected := 3
		got = cache.Get("3")

		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}

	})
}
