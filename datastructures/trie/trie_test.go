package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	n := NewNode()

	insertWords := [...]string{
		"nikola",
		"tesla",
	}
	checkWords := map[string]bool{
		"thomas": false,
		"edison": false,
		"nikola": true,
	}

	for _, w := range insertWords {
		n.Insert(w)
		t.Logf(
			"added \"%s\" to the Trie.",
			w,
		)
	}

	for k, v := range checkWords {
		ok := n.Find(k)
		if ok != v {
			t.Fatalf(
				"\"%s\" is supposed to be %sin the Trie.",
				k,
				map[bool]string{true: "", false: "NOT "}[v],
			)
		}
		t.Logf(
			"\"%s\" is %sin the Trie.",
			k,
			map[bool]string{true: "", false: "NOT "}[ok],
		)
	}
}

func BenchmarkTrie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := NewNode()

		n.Insert("nikola")
		n.Insert("tesla")

		n.Find("thomas")
		n.Find("edison")
		n.Find("nikola")
	}
}

func ExampleNode() {
	// creates a new node
	node := NewNode()

	// adds words
	node.Insert("nikola")
	node.Insert("tesla")

	// finds words
	node.Find("thomas") // false
	node.Find("edison") // false
	node.Find("nikola") // true
}
