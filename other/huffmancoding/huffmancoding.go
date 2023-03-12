// huffman.go
// description: Implements Huffman compression, encoding and decoding
// details:
// We implement the linear-time 2-queue method described here https://en.wikipedia.org/wiki/Huffman_coding.
// It assumes that the list of symbol-frequencies is sorted.
// author(s) [pedromsrocha](https://github.com/pedromsrocha)
// see also huffmancoding_test.go

// Package huffman provides primitives for doing Huffman compression, encodding and decoding.
package huffman

// A Node of an Huffman tree, which can either be a leaf or an internal node.
// Each node has a weight.
// A leaf node has an associated symbol, but no children (i.e., left == right == nil).
// A parent node has a left and right child and no symbol (i.e., symbol == -1).
type Node struct {
	left   *Node
	right  *Node
	symbol rune
	weight int
}

// A SymbolFreq is a pair of a symbol and its associated frequency.
type SymbolFreq struct {
	symbol rune
	freq   int
}

// The algorithm operates on []SymbolFreq sorted by frequency
type ByFreq []SymbolFreq

func (x ByFreq) Len() int           { return len(x) }
func (x ByFreq) Less(i, j int) bool { return x[i].freq < x[j].freq }
func (x ByFreq) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// BuildTree returns the root Node of the Huffman tree by compressing listfreq.
// The compression produces the most optimal code lengths, provided listfreq is ordered,
// i.e.: listfreq[i] <= listfreq[j], whenever i < j.
func BuildTree(listfreq ByFreq) Node {
	if len(listfreq) < 1 {
		panic("huffman: calling BuildTree with an empty list of symbol-frequency pairs")
	}
	q1 := make([]Node, len(listfreq))
	q2 := make([]Node, 0, len(listfreq))
	for i, x := range listfreq { // after the loop, q1 is a slice of leaf nodes representing listfreq
		q1[i] = Node{left: nil, right: nil, symbol: x.symbol, weight: x.freq}
	}
	//loop invariant: q1, q2 are ordered by increasing weights
	for len(q1)+len(q2) > 1 {
		var node1, node2 Node
		node1, q1, q2 = least(q1, q2)
		node2, q1, q2 = least(q1, q2)
		node := Node{left: &node1, right: &node2,
			symbol: -1, weight: node1.weight + node2.weight}
		q2 = append(q2, node)
	}
	if len(q1) == 1 { // returns the remaining node in q1, q2
		return q1[0]
	}
	return q2[0]
}

// least removes the node with lowest weight from q1, q2.
// It returns the node with lowest weight and the slices q1, q2 after the update.
func least(q1 []Node, q2 []Node) (Node, []Node, []Node) {
	if len(q1) == 0 {
		return q2[0], q1, q2[1:]
	}
	if len(q2) == 0 {
		return q1[0], q1[1:], q2
	}
	if q1[0].weight <= q2[0].weight {
		return q1[0], q1[1:], q2
	}
	return q2[0], q1, q2[1:]
}

// BuildDict recursively traverses the Huffman tree pointed by node to build dict, that maps symbols to binary codes,
// each code being prefixed by prefix.
// left and right children are labelled with the booleans false and true, respectively.
func BuildDict(node *Node, prefix []bool, dict map[rune][]bool) {
	if node.symbol != -1 { //base case
		dict[node.symbol] = prefix
	} else { // inductive step
		prefixLeft := make([]bool, len(prefix))
		copy(prefixLeft, prefix)
		prefixLeft = append(prefixLeft, false)
		BuildDict(node.left, prefixLeft, dict)
		prefixRight := make([]bool, len(prefix))
		copy(prefixRight, prefix)
		prefixRight = append(prefixRight, true)
		BuildDict(node.right, prefixRight, dict)
	}
}

// Encode encodes the string in using dictionary dict.
func Encode(dict map[rune][]bool, in string) []bool {
	out := make([]bool, 0)
	for _, s := range in {
		out = append(out, dict[s]...)
	}
	return out
}

// Decode recursively decodes the binary code in, by traversing the Huffman compression tree pointed by root.
// current stores the current node of the traversing algorithm.
// out stores the current decoded string.
func Decode(root, current *Node, in []bool, out string) string {
	if current.symbol != -1 {
		out += string(current.symbol)
		return Decode(root, root, in, out)
	}
	if len(in) == 0 {
		return out
	}
	if in[0] {
		return Decode(root, current.right, in[1:], out)
	} else {
		return Decode(root, current.left, in[1:], out)
	}
}
