// huffmancoding_test.go
// description: Tests the compresssion, encoding and decoding algorithms of huffmancoding.go.
// author(s) [pedromsrocha](https://github.com/pedromsrocha)
// see huffmancoding.go

package huffman

import (
	"testing"
)

func TestHuffman(t *testing.T) {
	var tests = []struct {
		message  string
		listfreq []SymbolFreq
	}{
		{"hello world \U0001F600",
			[]SymbolFreq{
				{'h', 1},
				{'e', 1},
				{'w', 1},
				{'r', 1},
				{'d', 1},
				{' ', 1},
				{'\U0001F600', 1},
				{'\n', 1},
				{'o', 2},
				{'l', 3}}},
	}
	for _, test := range tests {
		t.Run("huffman: "+test.message, func(t *testing.T) {
			tree := BuildTree(test.listfreq)
			dict := make(map[rune][]bool)
			BuildDict(&tree, nil, dict)
			messageCoded := Encode(dict, test.message)
			messageDecoded := Decode(&tree, &tree, messageCoded, "")
			if messageDecoded != test.message {
				t.Errorf("got: %v but expected: %v", messageDecoded, test.message)

			}
		})
	}
}
