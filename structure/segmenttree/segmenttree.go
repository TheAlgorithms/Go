//Segment Tree Data Structure for Range Queries
//Build: O(n*log(n))
//Query: O(log(n))
//reference: https://cp-algorithms.com/data_structures/segment_tree.html

package segmenttree

import (
	"github.com/TheAlgorithms/Go/dynamic"
)

const emptyLazyNode = -1

//SegmentTree with original array and the Segment Tree array
type SegmentTree struct {
	array       []int
	segmentTree []int
	lazyTree    []int
}

//Propagate lazy tree node values
func (s *SegmentTree) Propagate(node int, leftNode int, rightNode int) {
	if s.lazyTree[node] != emptyLazyNode {
		//add lazy node value multiplied by (right-left+1), which represents all interval
		//this is the same of adding a value on each node
		s.segmentTree[node] += (rightNode - leftNode + 1) * s.lazyTree[node]

		if leftNode == rightNode {
			//leaf node
			s.array[leftNode] = s.lazyTree[node]
		} else {
			//propagate lazy node value for children nodes
			s.lazyTree[2*node] = s.lazyTree[node]
			s.lazyTree[2*node+1] = s.lazyTree[node]
		}

		//clear lazy node
		s.lazyTree[node] = emptyLazyNode
	}
}

//Query on interval [firstIndex, leftIndex]
//node, leftNode and rightNode always should start with 1, 0 and len(array)-1
func (s *SegmentTree) Query(node int, leftNode int, rightNode int, firstIndex int, lastIndex int) int {
	if (firstIndex > lastIndex) || (leftNode > rightNode) {
		//outside the interval
		return 0
	}

	//propagate lazy tree
	s.Propagate(node, leftNode, rightNode)

	if (leftNode >= firstIndex) && (rightNode <= lastIndex) {
		//inside the interval
		return s.segmentTree[node]
	}

	//get sum of left and right nodes
	mid := (leftNode + rightNode) / 2

	leftNodeSum := s.Query(2*node, leftNode, mid, firstIndex, dynamic.Min(mid, lastIndex))
	rightNodeSum := s.Query(2*node+1, mid+1, rightNode, dynamic.Max(firstIndex, mid+1), lastIndex)

	return leftNodeSum + rightNodeSum
}

//Update Segment Tree
//node, leftNode and rightNode always should start with 1, 0 and len(array)-1
//index is the array index that you want to update
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
		s.lazyTree[node] = value
		s.Propagate(node, leftNode, rightNode)
	} else {
		//update left and right nodes
		mid := (leftNode + rightNode) / 2

		s.Update(2*node, leftNode, mid, firstIndex, dynamic.Min(mid, lastIndex), value)
		s.Update(2*node+1, mid+1, rightNode, dynamic.Max(firstIndex, mid+1), lastIndex, value)

		s.segmentTree[node] = s.segmentTree[2*node] + s.segmentTree[2*node+1]
	}
}

//Build Segment Tree
//node, leftNode and rightNode always should start with 1, 0 and len(array)-1
func (s *SegmentTree) Build(node int, left int, right int) {
	if left == right {
		//leaf node
		s.segmentTree[node] = s.array[left]
	} else {
		//get sum of left and right nodes
		mid := (left + right) / 2

		s.Build(2*node, left, mid)
		s.Build(2*node+1, mid+1, right)

		s.segmentTree[node] = s.segmentTree[2*node] + s.segmentTree[2*node+1]
	}
}

func NewSegmentTree(array []int) *SegmentTree {
	if len(array) == 0 {
		return nil
	}

	segTree := SegmentTree{
		array:       array,
		segmentTree: make([]int, 4*len(array)),
		lazyTree:    make([]int, 4*len(array)),
	}

	for i := range segTree.lazyTree {
		//fill lazyTree with empty lazy nodes
		segTree.lazyTree[i] = emptyLazyNode
	}

	//starts with node 1 and interval [0, len(arr)-1] inclusive
	segTree.Build(1, 0, len(array)-1)

	return &segTree
}
