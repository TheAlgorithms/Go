package hashmap

import (
	"testing"
)

func TestHashTable(t *testing.T) {

	mp := newTable()

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

	t.Run("Test 6: Checking if the key that doesnt exists returns false", func(t *testing.T) {
		want := false
		got := mp.Contains(2)
		if got != want {
			t.Errorf("Key '2' doesn't exists in the map but it says otherwise")
		}
	})
	t.Run("Test 7: Checking deleted key exists or not", func(t *testing.T) {
		want := false
		mp.Delete(2)
		got := mp.Contains(2)
		if got != want {
			t.Errorf("Key '2' deleted from table but it says otherwise")
		}
	})
	t.Run("Test 8: HashTable Implement table interface", func(t *testing.T) {
		want := true
		var l interface{} = newTable()
		_, got := l.(HashTable)
		if got != want {
			t.Errorf("table should implement HashTable interface but it say otherwise")
		}
	})

}

// Put Benchmark on HashTable with Collision Handled via openaddress
func BenchmarkOpenAddresPut(b *testing.B) {
	tb := newTable()
	for i := 0; i < b.N; i++ {
		tb.Put(i, i)
	}
}

// Go map Put benchmark
func BenchmarkGoPut(b *testing.B) {
	tb := make(map[interface{}]interface{}, 2)
	for i := 0; i < b.N; i++ {
		tb[i] = i
	}
}

// Put Benchmark on HashTable with Collision Handled via Chaining
func BenchmarkOldPut(b *testing.B) {
	tb := New()
	for i := 0; i < b.N; i++ {
		tb.Put(i, i)
	}
}
