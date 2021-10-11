//Segment Tree Data Structure for Range Queries
//Build: O(n*log(n))
//Query: O(log(n))
//reference: https://cp-algorithms.com/data_structures/segment_tree.html

package segmenttree

import "github.com/TheAlgorithms/Go/dynamic"

//SegmentTree with original array and the Segment Tree array
type SegmentTree struct {
	array       []int
	segmentTree []int
}

//Query on interval [firstIndex, leftIndex]
//node, leftNode and rightNode always should start with 1, 0 and len(array)-1
func (s *SegmentTree) Query(node int, leftNode int, rightNode int, firstIndex int, lastIndex int) int {
	if (firstIndex > lastIndex) || (leftNode > rightNode) {
		//outside the interval
		return 0
	}

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
	}

	//starts with node 1 and interval [0, len(arr)-1] inclusive
	segTree.Build(1, 0, len(array)-1)

	return &segTree
}
