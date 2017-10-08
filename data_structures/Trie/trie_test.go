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
		n.insert(w)
		t.Logf(
			"added \"%s\" to the Trie.",
			w,
		)
	}

	for k, v := range checkWords {
		ok := n.find(k)
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

		n.insert("nikola")
		n.insert("tesla")

		n.find("thomas")
		n.find("edison")
		n.find("nikola")
	}
}

func ExampleNode() {
	// creates a new node
	node := NewNode()

	// adds words
	node.insert("nikola")
	node.insert("tesla")

	// finds words
	node.find("thomas") // false
	node.find("edison") // false
	node.find("nikola") // true
}
