// huffman.go
// description: Implements Huffman compression, associated encoding and decoding
// author(s) [pedromsrocha](https://github.com/pedromsrocha)
// reference: https://en.wikipedia.org/wiki/Huffman_coding

package huffman

type Node struct {
	left   *Node
	right  *Node
	symbol rune
	weight int
}

type SymbolFreq struct {
	symbol rune
	freq   int
}

func BuildTree(listfreq []SymbolFreq) Node {
	q1 := make([]Node, len(listfreq))
	q2 := make([]Node, 0, len(listfreq))
	for i, x := range listfreq {
		q1[i] = Node{left: nil, right: nil, symbol: x.symbol, weight: x.freq}
	}
	for len(q1)+len(q2) > 1 {
		var node1, node2 Node
		node1, q1, q2 = Pop(q1, q2)
		node2, q1, q2 = Pop(q1, q2)
		node := Node{left: &node1, right: &node2,
			symbol: -1, weight: node1.weight + node2.weight}
		q2 = append(q2, node)
	}
	node, _, _ := Pop(q1, q2)
	return node
}

func Pop(q1 []Node, q2 []Node) (Node, []Node, []Node) {
	if len(q1) == 0 {
		return q2[0], q1, q2[1:]
	}
	if len(q2) == 0 {
		return q1[0], q1[1:], q2
	}
	if q1[0].weight <= q2[0].weight {
		return q1[0], q1[1:], q2
	} else {
		return q2[0], q1, q2[1:]
	}
}

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

func Encode(dict map[rune][]bool, in string) []bool {
	out := make([]bool, 0)
	for _, s := range in {
		out = append(out, dict[s]...)
	}
	return out
}

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
