package trie

import (
	"testing"
)

func TestTrieInsert(t *testing.T) {
	n := NewNode()

	insertWords := []string{
		"nikola",
		"tesla",
	}

	checkWords := map[string]bool{
		"thomas": false,
		"edison": false,
		"nikola": true,
	}

	n.Insert(insertWords...)
	n.verify(t, checkWords)
	n.verifySizeCapa(t, 2, 12)

}

func TestTrieRemove(t *testing.T) {
	n := NewNode()

	insertWords := []string{
		"nikola",
		"tesla",
		"albert",
		"einstein",
	}

	checkWords := map[string]bool{
		"thomas":   false,
		"edison":   false,
		"nikola":   true,
		"albert":   true,
		"einstein": true,
	}

	n.Insert(insertWords...)
	n.verify(t, checkWords)
	size, capa := n.Size(), n.Capacity() // 4 words ...

	n.Remove("albert")
	checkWords["albert"] = false
	size-- // 3 words in size, but no change in capacity
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa)

	n.Remove("albert") // no effect since already removed
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa)

	n.Remove("marcel") // no effect since ,o, existant
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa)

	n.Remove("nikola", "tesla") // 1 word
	checkWords["nikola"] = false
	checkWords["tesla"] = false
	size -= 2 // t1 word left,  but still no change in capacity
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa)

	// compact the Tree
	if n.Compact() {
		t.Fatal("the Trie should not be completely removable after compaction")
	}
	if capa == n.Capacity() {
		t.Fatal("capacity should have reduced following compaction")
	}
	capa = n.Capacity()
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa) // still 1 word, reduced capacity

	n.Remove("einstein")
	checkWords["einstein"] = false
	size-- // No more words
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa) // no words, but still have some nodes left capacity

	if !n.Compact() {
		t.Fatal("the root node of an empty Trie should be marked as removable after compaction")
	}
	n.verifySizeCapa(t, 0, 1) // no words, only the root node left
}

// --------------- helper functions ---------------------------

// verify if provided words are present
func (n *Node) verify(t *testing.T, checkWords map[string]bool) {
	for k, v := range checkWords {
		ok := n.Find(k)
		if ok != v {
			t.Fatalf(
				"%q is %s supposed to be in the Trie.",
				k,
				map[bool]string{true: "", false: "NOT "}[v],
			)
		}
		// t.Logf(
		// 	"\"%s\" is %sin the Trie.",
		// 	k,
		// 	map[bool]string{true: "", false: "NOT "}[ok],
		// )
	}
}

// verify expected size and capacity
func (n *Node) verifySizeCapa(t *testing.T, expectedSize, expectedCapacity int) {
	if got := n.Size(); got != expectedSize {
		t.Fatalf("Expected Size was %d but got %d", expectedSize, got)
	}
	if got := n.Capacity(); got != expectedCapacity {
		t.Fatalf("Expected Capacity was %d but got %d", expectedCapacity, got)
	}
}
