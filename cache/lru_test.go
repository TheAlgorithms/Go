package cache_test

import (
	"testing"

	lru "github.com/TheAlgorithms/Go/cache"
)

func TestLRU(t *testing.T) {

	cache := lru.New(3)

	t.Run("Test 1: Put number and Get is correct", func(t *testing.T) {
		key, value := "1", 1
		cache.Put(key, value)
		got := cache.Get(key)

		if got != value {
			t.Errorf("Put: %v, Got: %v", value, got)
		}
	})

	t.Run("Test 2: Get data without cache should got nil", func(t *testing.T) {
		got := cache.Get("2")
		if got != nil {
			t.Errorf("Got: %v", got)
		}
	})

	t.Run("Test 3: Put data over capacity and Get should got nil", func(t *testing.T) {
		cache.Put("2", 2)
		cache.Put("3", 3)
		cache.Put("4", 4)

		got := cache.Get("1")
		if got != nil {
			t.Errorf("Got: %v", got)
		}
	})

	t.Run("test 4: Put key over capacity but key already existing", func(t *testing.T) {
		cache.Put("4", 4)
		cache.Put("3", 3)
		cache.Put("2", 2)
		cache.Put("1", 1)

		got := cache.Get("4")
		if got != nil {
			t.Errorf("Got: %v", got)
		}

		expected := 3
		got = cache.Get("3")

		if got != expected {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}

	})
}
