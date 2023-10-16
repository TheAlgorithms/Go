// Package sqrt contains algorithms and data structures that contains a √n in their complexity
package sqrt

import "math"

// Sqrt (or Square Root) Decomposition is a technique used for query an array and perform updates
// Inside this package is described its most simple data structure, you can find more at: https://cp-algorithms.com/data_structures/sqrt_decomposition.html
//
// Formally, You can use SqrtDecomposition only if:
//
// Given a function $Query:E_1,...,E_n\rightarrow Q$
//
// if $\exist unionQ:Q,Q\rightarrow Q$
//
// s.t.
//
// - $\forall n\in \N > 1, 1\le i<n, E_1,..., E_n\in E \\ query(E_1,..., E_n)=unionQ(query(E_1,..., E_i), query(E_{i+1},...,E_n))$
//
// - (Only if you want use $update$ function)
// $\forall n\in \N > 0, E_1,..., E_n\in E \\ query(E_1,...,E_{new},..., E_n)=updateQ(query(E_1,...,E_{old},...,E_n), indexof(E_{old}), E_{new})$
type SqrtDecomposition[E any, Q any] struct {
	querySingleElement func(element E) Q
	unionQ             func(q1 Q, q2 Q) Q
	updateQ            func(oldQ Q, oldE E, newE E) (newQ Q)

	elements  []E
	blocks    []Q
	blockSize uint64
}

// Create a new SqrtDecomposition instance with the parameters as specified by SqrtDecomposition comment
// Assumptions:
//   - len(elements) > 0
func NewSqrtDecomposition[E any, Q any](
	elements []E,
	querySingleElement func(element E) Q,
	unionQ func(q1 Q, q2 Q) Q,
	updateQ func(oldQ Q, oldE E, newE E) (newQ Q),
) *SqrtDecomposition[E, Q] {
	sqrtDec := &SqrtDecomposition[E, Q]{
		querySingleElement: querySingleElement,
		unionQ:             unionQ,
		updateQ:            updateQ,
		elements:           elements,
	}
	sqrt := math.Sqrt(float64(len(sqrtDec.elements)))
	blockSize := uint64(sqrt)
	numBlocks := uint64(math.Ceil(float64(len(elements)) / float64(blockSize)))
	sqrtDec.blocks = make([]Q, numBlocks)
	for i := uint64(0); i < uint64(len(elements)); i++ {
		if i%blockSize == 0 {
			sqrtDec.blocks[i/blockSize] = sqrtDec.querySingleElement(elements[i])
		} else {
			sqrtDec.blocks[i/blockSize] = sqrtDec.unionQ(sqrtDec.blocks[i/blockSize], sqrtDec.querySingleElement(elements[i]))
		}
	}
	sqrtDec.blockSize = blockSize
	return sqrtDec
}

// Performs a query from index start to index end (non included)
// Assumptions:
//   - start < end
//   - start and end are valid
func (s *SqrtDecomposition[E, Q]) Query(start uint64, end uint64) Q {
	firstIndexNextBlock := ((start / s.blockSize) + 1) * s.blockSize
	q := s.querySingleElement(s.elements[start])
	if firstIndexNextBlock > end { // if in same block
		start++
		for start < end {
			q = s.unionQ(q, s.querySingleElement(s.elements[start]))
			start++
		}
	} else {
		// left side
		start++
		for start < firstIndexNextBlock {
			q = s.unionQ(q, s.querySingleElement(s.elements[start]))
			start++
		}

		//middle part
		endBlock := end / s.blockSize
		for i := firstIndexNextBlock / s.blockSize; i < endBlock; i++ {
			q = s.unionQ(q, s.blocks[i])
		}

		// right part
		for i := endBlock * s.blockSize; i < end; i++ {
			q = s.unionQ(q, s.querySingleElement(s.elements[i]))
		}
	}
	return q
}

// Assumptions:
//   - index is valid
func (s *SqrtDecomposition[E, Q]) Update(index uint64, newElement E) {
	i := index / s.blockSize
	s.blocks[i] = s.updateQ(s.blocks[i], s.elements[index], newElement)
	s.elements[index] = newElement
}
