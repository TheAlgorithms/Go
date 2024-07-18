package hashmap_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/structure/hashmap"
)

func TestHashMap(t *testing.T) {
	mp := hashmap.DefaultNew()

	t.Run("Test 1: Put(10) and checking if Get() is correct", func(t *testing.T) {
		mp.Put("test", 10)
		got := mp.Get("test")
		if got != 10 {
			t.Errorf("Put: %v, Got: %v", 10, got)
		}
	})

	t.Run("Test 2: Reassiging the value and checking if its correct", func(t *testing.T) {
		mp.Put("test", 20)
		got := mp.Get("test")
		if got != 20 {
			t.Errorf("Put (reassign): %v, Got: %v", 20, got)
		}
	})

	t.Run("Test 3: Adding new key when there is already some data", func(t *testing.T) {
		mp.Put("test2", 30)
		got := mp.Get("test2")
		if got != 30 {
			t.Errorf("Put: %v, Got: %v", got, 30)
		}
	})

	t.Run("Test 4: Adding numeric key", func(t *testing.T) {
		mp.Put(1, 40)
		got := mp.Get(1)
		if got != 40 {
			t.Errorf("Put: %v, Got: %v", got, 40)
		}
	})

	t.Run("Test 5: Checking the Contains function", func(t *testing.T) {
		want := true
		got := mp.Contains(1)
		if want != got {
			t.Errorf("Key '1' exists but couldn't be retrieved")
		}
	})

	t.Run("Test 6: Checking if the key that does not exist returns false", func(t *testing.T) {
		want := false
		got := mp.Contains(2)
		if got != want {
			t.Errorf("Key '2' does not exist in the map but it says otherwise")
		}
	})

	t.Run("Test 7: Checking if the key does not exist Get func returns nil", func(t *testing.T) {
		want := any(nil)
		got := mp.Get(2)
		if got != want {
			t.Errorf("Key '2' does not exists in the map but it says otherwise")
		}
	})

	t.Run("Test 8: Resizing a map", func(t *testing.T) {
		mp := hashmap.New(4, 4)

		for i := 0; i < 20; i++ {
			mp.Put(i, 40)
		}

		got := mp.Get(5)
		if got != 40 {
			t.Errorf("Put: %v, Got: %v", got, 40)
		}
	})
}
