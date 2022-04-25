//Segment Tree Data Structure for Range Queries
//Build: O(n*log(n))
//Query: O(log(n))
//Update: O(log(n))
//reference: https://cp-algorithms.com/data_structures/segment_tree.html

package segmenttree

import (
	"github.com/TheAlgorithms/Go/math/max"
	"github.com/TheAlgorithms/Go/math/min"
)

const emptyLazyNode = 0

//SegmentTree with original Array and the Segment Tree Array
type SegmentTree struct {
	Array       []int
	SegmentTree []int
	LazyTree    []int
}

//Propagate lazy tree node values
func (s *SegmentTree) Propagate(node int, leftNode int, rightNode int) {
	if s.LazyTree[node] != emptyLazyNode {
		//add lazy node value multiplied by (right-left+1), which represents all interval
		//this is the same of adding a value on each node
		s.SegmentTree[node] += (rightNode - leftNode + 1) * s.LazyTree[node]

		if leftNode == rightNode {
			//leaf node
			s.Array[leftNode] += s.LazyTree[node]
		} else {
			//propagate lazy node value for children nodes
			//may propagate multiple times, children nodes should accumulate lazy node value
			s.LazyTree[2*node] += s.LazyTree[node]
			s.LazyTree[2*node+1] += s.LazyTree[node]
		}

		//clear lazy node
		s.LazyTree[node] = emptyLazyNode
	}
}

//Query on interval [firstIndex, leftIndex]
//node, leftNode and rightNode always should start with 1, 0 and len(Array)-1
func (s *SegmentTree) Query(node int, leftNode int, rightNode int, firstIndex int, lastIndex int) int {
	if (firstIndex > lastIndex) || (leftNode > rightNode) {
		//outside the interval
		return 0
	}

	//propagate lazy tree
	s.Propagate(node, leftNode, rightNode)

	if (leftNode >= firstIndex) && (rightNode <= lastIndex) {
		//inside the interval
		return s.SegmentTree[node]
	}

	//get sum of left and right nodes
	mid := (leftNode + rightNode) / 2

	leftNodeSum := s.Query(2*node, leftNode, mid, firstIndex, min.Int(mid, lastIndex))
	rightNodeSum := s.Query(2*node+1, mid+1, rightNode, max.Int(firstIndex, mid+1), lastIndex)

	return leftNodeSum + rightNodeSum
}

//Update Segment Tree
//node, leftNode and rightNode always should start with 1, 0 and len(Array)-1
//index is the Array index that you want to update
//value is the value that you want to override
func (s *SegmentTree) Update(node int, leftNode int, rightNode int, firstIndex int, lastIndex int, value int) {
	//propagate lazy tree
	s.Propagate(node, leftNode, rightNode)

	if (firstIndex > lastIndex) || (leftNode > rightNode) {
		//outside the interval
		return
	}

	if (leftNode >= firstIndex) && (rightNode <= lastIndex) {
		//inside the interval
		//accumulate the lazy node value
		s.LazyTree[node] += value
		s.Propagate(node, leftNode, rightNode)
	} else {
		//update left and right nodes
		mid := (leftNode + rightNode) / 2

		s.Update(2*node, leftNode, mid, firstIndex, min.Int(mid, lastIndex), value)
		s.Update(2*node+1, mid+1, rightNode, max.Int(firstIndex, mid+1), lastIndex, value)

		s.SegmentTree[node] = s.SegmentTree[2*node] + s.SegmentTree[2*node+1]
	}
}

//Build Segment Tree
//node, leftNode and rightNode always should start with 1, 0 and len(Array)-1
func (s *SegmentTree) Build(node int, left int, right int) {
	if left == right {
		//leaf node
		s.SegmentTree[node] = s.Array[left]
	} else {
		//get sum of left and right nodes
		mid := (left + right) / 2

		s.Build(2*node, left, mid)
		s.Build(2*node+1, mid+1, right)

		s.SegmentTree[node] = s.SegmentTree[2*node] + s.SegmentTree[2*node+1]
	}
}

func NewSegmentTree(Array []int) *SegmentTree {
	if len(Array) == 0 {
		return nil
	}

	segTree := SegmentTree{
		Array:       Array,
		SegmentTree: make([]int, 4*len(Array)),
		LazyTree:    make([]int, 4*len(Array)),
	}

	for i := range segTree.LazyTree {
		//fill LazyTree with empty lazy nodes
		segTree.LazyTree[i] = emptyLazyNode
	}

	//starts with node 1 and interval [0, len(arr)-1] inclusive
	segTree.Build(1, 0, len(Array)-1)

	return &segTree
}
