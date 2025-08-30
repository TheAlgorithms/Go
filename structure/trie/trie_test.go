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

func TestTrieInsert_substrings(t *testing.T) {
	n := NewNode()

	insertWords := []string{
		"aa",
		"aaaa",
		"aaaaa",
	}

	checkWords := map[string]bool{
		"a":       false,
		"aa":      true,
		"aaa":     false,
		"aaaa":    true,
		"aaaaa":   true,
		"aaaaaa":  false,
		"aaaaaaa": false,
	}

	n.Insert(insertWords...)
	n.verify(t, checkWords)
	n.verifySizeCapa(t, 3, 5+1)

	n.Remove("aaaa")
	checkWords["aaaa"] = false
	n.verify(t, checkWords)
	n.verifySizeCapa(t, 2, 5+1)

	if n.Compact() {
		t.Fatalf("it should not be possible to remove the node")
	}
	n.verify(t, checkWords)
	n.verifySizeCapa(t, 2, 5+1)
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

	n.Remove("marcel") // no effect since ,o, existent
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
	if capa <= n.Capacity() {
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

// Test inserting empty string
func TestTrieInsertEmptyString(t *testing.T) {
	root := NewNode()
	root.Insert("")

	check := map[string]bool{
		"": false,
	}
	root.verify(t, check)
	root.verifySizeCapa(t, 0, root.Capacity())
}

// Test inserting words with shared prefix
func TestTrieInsertSharedPrefix(t *testing.T) {
	root := NewNode()
	root.Insert("abc", "abcd", "abcde")

	check := map[string]bool{
		"abc":   true,
		"abcd":  true,
		"abcde": true,
		"ab":    false,
	}
	root.verify(t, check)

	root.verifySizeCapa(t, 3, root.Capacity())
}

// Test removing non-existent words
func TestTrieRemoveNonExistentEdge(t *testing.T) {
	root := NewNode()
	root.Insert("abc", "def")

	root.Remove("xyz") // does not exist

	check := map[string]bool{
		"abc": true,
		"def": true,
		"xyz": false,
	}
	root.verify(t, check)
	root.verifySizeCapa(t, 2, root.Capacity())
}

// Test inserting duplicates multiple times
func TestTrieInsertDuplicatesMultiple(t *testing.T) {
	root := NewNode()
	root.Insert("abc", "abc", "abc")

	check := map[string]bool{
		"abc": true,
	}
	root.verify(t, check)

	// Size should count only once
	root.verifySizeCapa(t, 1, root.Capacity())
}

// Test removing all words sequentially
func TestTrieRemoveAllSequential(t *testing.T) {
	root := NewNode()
	root.Insert("a", "b", "c")

	root.Remove("a")
	root.Remove("b")
	root.Remove("c")

	check := map[string]bool{
		"a": false,
		"b": false,
		"c": false,
	}
	root.verify(t, check)

	// Size should be zero
	root.verifySizeCapa(t, 0, root.Capacity())
}

// Test inserting substrings and removing middle one
func TestTrieInsertSubstringsRemoveMiddle(t *testing.T) {
	root := NewNode()
	root.Insert("a", "ab", "abc")

	root.Remove("ab")

	check := map[string]bool{
		"a":   true,
		"ab":  false,
		"abc": true,
	}
	root.verify(t, check)

	// Size should be 2
	root.verifySizeCapa(t, 2, root.Capacity())
}

// Test Compact on empty Trie
func TestTrieCompactEmptyTrie(t *testing.T) {
	root := NewNode()
	if !root.Compact() {
		t.Fatal("Empty Trie root should be removable after compaction")
	}
	root.verifySizeCapa(t, 0, 1) // only root remains
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
