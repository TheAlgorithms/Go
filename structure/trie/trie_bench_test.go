package trie

import (
	"fmt"
	"math/rand"
	"testing"
)

// CAUTION : make sure to limit the benchamrks to 3000 iterations,
// or removal will mostly process an empty Trie, giving absurd results.

// go test -v -bench=. -benchmem -benchtime=3000x

// === RUN   TestTrieInsert
// --- PASS: TestTrieInsert (0.00s)
// === RUN   TestTrieRemove
// --- PASS: TestTrieRemove (0.00s)
// === RUN   ExampleNode
// --- PASS: ExampleNode (0.00s)
// goos: linux
// goarch: amd64
// pkg: github.com/TheAlgorithms/Go/structure/trie
// BenchmarkTrie_Insert
// BenchmarkTrie_Insert-8                      3000           1559881 ns/op         1370334 B/op      25794 allocs/op
// BenchmarkTrie_Find_non_existant
// BenchmarkTrie_Find_non_existant-8           3000                59.1 ns/op             0 B/op          0 allocs/op
// BenchmarkTrie_Find_existant
// BenchmarkTrie_Find_existant-8               3000               238 ns/op               0 B/op          0 allocs/op
// BenchmarkTrie_Remove_lazy
// BenchmarkTrie_Remove_lazy-8                 3000               126 ns/op               0 B/op          0 allocs/op
// BenchmarkTrie_Remove_and_Compact
// BenchmarkTrie_Remove_and_Compact-8          3000            213945 ns/op               0 B/op          0 allocs/op
// PASS
// ok      github.com/TheAlgorithms/Go/structure/trie      5.355s

func BenchmarkTrie_Insert(b *testing.B) {

	// prepare random words for insertion
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}

	// Reset timer and run benchmark for insertion
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := NewNode()
		n.Insert(insert...)
	}
}

func BenchmarkTrie_Find_non_existant(b *testing.B) {

	// prepare random words for insertion
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	// Reset timer and run benchmark for finding non existing
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Find("0.3213213244346546546546565465465") // does not exists
	}
}

func BenchmarkTrie_Find_existant(b *testing.B) {
	// prepare and insert random words
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	// Reset timer and run benchmark for finding existing words
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Find(insert[i%3000]) // always exists !
	}
}

func BenchmarkTrie_Remove_lazy(b *testing.B) {
	// prepare and insert random words
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	// Reset timer and run benchmark for lazily removing words
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Remove(insert[i%3000]) //  exists, at least until removed ...
	}
}

func BenchmarkTrie_Remove_and_Compact(b *testing.B) {
	// prepare and insert random words
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	// Reset timer and run benchmark for removing words and immediately compacting
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Remove(insert[i%3000])
		n.Compact()
	}
}
