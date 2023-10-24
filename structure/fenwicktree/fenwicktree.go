//Fenwick Tree Data Structure for Range Queries. Also known as Binary Indexed Tree
//Build: O(N)
//Query: O(log(N))
//Update: O(log(N))
//reference: https://brilliant.org/wiki/fenwick-tree/

package fenwicktree

// n: Size of the input array.
// array: the input array on which queries are made.
// bit: store the sum of ranges.
type FenwickTree struct {
	n     int
	array []int
	bit   []int
}

// This function creates a new Fenwick tree, initialized array and bit with
// the values of the array. Note that the queries and updates should have
// one based indexing.
func NewFenwickTree(array []int) *FenwickTree {
	newArray := []int{0} // Appending a 0 to the beginning as this implementation uses 1 based indexing
	fenwickTree := &FenwickTree{
		n:     len(array),
		array: append(newArray, array...),
		bit:   append(newArray, array...),
	}
	for i := 1; i < fenwickTree.n; i++ {
		nextPos := i + (i & -i)
		if nextPos <= fenwickTree.n {
			fenwickTree.bit[nextPos] += fenwickTree.bit[i]
		}
	}
	return fenwickTree
}

// Returns the sum of the prefix ending at position pos.
func (f *FenwickTree) PrefixSum(pos int) int {
	if pos > f.n {
		return 0
	}
	prefixSum := 0
	for i := pos; i > 0; i -= (i & -i) {
		prefixSum += f.bit[i]
	}
	return prefixSum
}

// Returns the sum of the elements in the range l to r
// both inclusive.
func (f *FenwickTree) RangeSum(l int, r int) int {
	return f.PrefixSum(r) - f.PrefixSum(l-1)
}

// Adds value to the element at position pos of the array
// and recomputes the range sums.
func (f *FenwickTree) Add(pos int, value int) {
	for i := pos; i <= f.n; i += (i & -i) {
		f.bit[i] += value
	}
}
