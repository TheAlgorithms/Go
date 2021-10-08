//Segment Tree Data Structure for Range Queries
//Build: O(n*log(n))
//Query: O(log(n))
//reference: https://cp-algorithms.com/data_structures/segment_tree.html

package segmenttree

//SegmentTree with original array and the Segment Tree array
type SegmentTree struct {
	array       []int
	segmentTree []int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//Query on interval [firstIndex, leftIndex]
//node, leftNode and rightNode always should start with 1, 0 and len(array)-1
func (s *SegmentTree) query(node int, leftNode int, rightNode int, firstIndex int, lastIndex int) int {
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

	leftNodeSum := s.query(2*node, leftNode, mid, firstIndex, min(mid, lastIndex))
	rightNodeSum := s.query(2*node+1, mid+1, rightNode, max(firstIndex, mid+1), lastIndex)

	return leftNodeSum + rightNodeSum
}

//Build Segment Tree
//node, leftNode and rightNode always should start with 1, 0 and len(array)-1
func (s *SegmentTree) build(node int, left int, right int) {
	if left == right {
		//leaf node
		s.segmentTree[node] = s.array[left]
	} else {
		//get sum of left and right nodes
		mid := (left + right) / 2

		s.build(2*node, left, mid)
		s.build(2*node+1, mid+1, right)

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
	segTree.build(1, 0, len(array)-1)

	return &segTree
}
