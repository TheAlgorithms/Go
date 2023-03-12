// huffmancoding_test.go
// description: Tests the compression, encoding and decoding algorithms of huffmancoding.go.
// author(s) [pedromsrocha](https://github.com/pedromsrocha)
// see huffmancoding.go

package huffman

import (
	"sort"
	"testing"
)

// The algorithm operates on []SymbolFreq sorted by frequency
type ByFreq []SymbolFreq

func (x ByFreq) Len() int           { return len(x) }
func (x ByFreq) Less(i, j int) bool { return x[i].freq < x[j].freq }
func (x ByFreq) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// SymbolCountOrd computes sorted symbol-frequency list of input message
func SymbolCountOrd(message string) ByFreq {
	runeCount := make(map[rune]int)
	for _, s := range message {
		runeCount[s]++
	}
	listfreq := make(ByFreq, len(runeCount))
	i := 0
	for s, n := range runeCount {
		listfreq[i] = SymbolFreq{symbol: s, freq: n}
		i++
	}
	sort.Sort(listfreq)
	return listfreq
}

func TestHuffman(t *testing.T) {
	messages := []string{
		"hello world \U0001F600",
		"colorless green ideas sleep furiously",
		"the quick brown fox jumps over the lazy dog",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
		Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
		aliquip ex ea commodo consequat.`,
	}

	for _, message := range messages {
		t.Run("huffman: "+message, func(t *testing.T) {
			tree, _ := BuildTree(SymbolCountOrd(message))
			codes := make(map[rune][]bool)
			GetEncoding(tree, nil, codes)
			messageCoded := Encode(codes, message)
			messageDecoded := Decode(tree, tree, messageCoded, "")
			if messageDecoded != message {
				t.Errorf("got: %q\nbut expected: %q", messageDecoded, message)

			}
		})
	}
}
