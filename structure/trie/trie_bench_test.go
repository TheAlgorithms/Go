package trie

import (
	"fmt"
	"math/rand"
	"testing"
)

// CAUTION : make sure to limit the benchmarks to 3000 iterations,
// or removal will mostly process an empty Trie, giving absurd results.

func BenchmarkTrie_Insert(b *testing.B) {
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := NewNode()
		n.Insert(insert...)
	}
}

func BenchmarkTrie_Find_non_existent(b *testing.B) {

	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Find("0.3213213244346546546546565465465") // does not exists
	}
}

func BenchmarkTrie_Find_existent(b *testing.B) {
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Find(insert[i%3000]) // always exists !
	}
}

func BenchmarkTrie_Remove_lazy(b *testing.B) {
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Remove(insert[i%3000]) //  exists, at least until removed ...
	}
}

func BenchmarkTrie_Remove_and_Compact(b *testing.B) {
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Remove(insert[i%3000])
		n.Compact()
	}
}
