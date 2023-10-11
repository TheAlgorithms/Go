// package set implements a Set using a golang map.
// This implies that only the types that are accepted as valid map keys can be used as set elements.
// For instance, do not try to Add a slice, or the program will panic.
package set

// New gives new set.
func New[T comparable](items ...T) Set[T] {
	st := set[T]{
		elements: make(map[T]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

// Set is an interface of possible methods on 'set'.
type Set[T comparable] interface {
	// Add: adds new element to the set
	Add(item T)
	// Delete: deletes the passed element from the set if present
	Delete(item T)
	// Len: gives the length of the set (total no. of elements in set)
	Len() int
	// GetItems: gives the array( []T ) of elements of the set.
	GetItems() []T
	// In: checks whether item is present in set or not.
	In(item T) bool
	// IsSubsetOf: checks whether set is subset of set2 or not.
	IsSubsetOf(set2 Set[T]) bool
	// IsProperSubsetOf: checks whether set is proper subset of set2 or not.
	// ex: [1,2,3] proper subset of [1,2,3,4] -> true
	IsProperSubsetOf(set2 Set[T]) bool
	// IsSupersetOf: checks whether set is superset of set2 or not.
	IsSupersetOf(set2 Set[T]) bool
	// IsProperSupersetOf: checks whether set is proper superset of set2 or not.
	// ex: [1,2,3,4] proper superset of [1,2,3] -> true
	IsProperSupersetOf(set2 Set[T]) bool
	// Union: gives new union set of both sets.
	// ex: [1,2,3] union [3,4,5] -> [1,2,3,4,5]
	Union(set2 Set[T]) Set[T]
	// Intersection: gives new intersection set of both sets.
	// ex: [1,2,3] Intersection [3,4,5] -> [3]
	Intersection(set2 Set[T]) Set[T]
	// Difference: gives new difference set of both sets.
	// ex: [1,2,3] Difference [3,4,5] -> [1,2]
	Difference(set2 Set[T]) Set[T]
	// SymmetricDifference: gives new symmetric difference set of both sets.
	// ex: [1,2,3] SymmetricDifference [3,4,5] -> [1,2,4,5]
	SymmetricDifference(set2 Set[T]) Set[T]
}

type set[T comparable] struct {
	elements map[T]bool
}

func (st *set[T]) Add(value T) {
	st.elements[value] = true
}

func (st *set[T]) Delete(value T) {
	delete(st.elements, value)
}

func (st *set[T]) GetItems() []T {
	keys := make([]T, 0, len(st.elements))
	for k := range st.elements {
		keys = append(keys, k)
	}
	return keys
}

func (st *set[T]) Len() int {
	return len(st.elements)
}

func (st *set[T]) In(value T) bool {
	if _, in := st.elements[value]; in {
		return true
	}
	return false
}

func (st *set[T]) IsSubsetOf(superSet Set[T]) bool {
	if st.Len() > superSet.Len() {
		return false
	}

	for _, item := range st.GetItems() {
		if !superSet.In(item) {
			return false
		}
	}
	return true
}

func (st *set[T]) IsProperSubsetOf(superSet Set[T]) bool {
	if st.Len() == superSet.Len() {
		return false
	}
	return st.IsSubsetOf(superSet)
}

func (st *set[T]) IsSupersetOf(subSet Set[T]) bool {
	return subSet.IsSubsetOf(st)
}

func (st *set[T]) IsProperSupersetOf(subSet Set[T]) bool {
	if st.Len() == subSet.Len() {
		return false
	}
	return st.IsSupersetOf(subSet)
}

func (st *set[T]) Union(st2 Set[T]) Set[T] {
	unionSet := New[T]()
	for _, item := range st.GetItems() {
		unionSet.Add(item)
	}
	for _, item := range st2.GetItems() {
		unionSet.Add(item)
	}
	return unionSet
}

func (st *set[T]) Intersection(st2 Set[T]) Set[T] {
	intersectionSet := New[T]()
	var minSet, maxSet Set[T]
	if st.Len() > st2.Len() {
		minSet = st2
		maxSet = st
	} else {
		minSet = st
		maxSet = st2
	}
	for _, item := range minSet.GetItems() {
		if maxSet.In(item) {
			intersectionSet.Add(item)
		}
	}
	return intersectionSet
}

func (st *set[T]) Difference(st2 Set[T]) Set[T] {
	differenceSet := New[T]()
	for _, item := range st.GetItems() {
		if !st2.In(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

func (st *set[T]) SymmetricDifference(st2 Set[T]) Set[T] {
	symmetricDifferenceSet := New[T]()
	dropSet := New[T]()
	for _, item := range st.GetItems() {
		if st2.In(item) {
			dropSet.Add(item)
		} else {
			symmetricDifferenceSet.Add(item)
		}
	}
	for _, item := range st2.GetItems() {
		if !dropSet.In(item) {
			symmetricDifferenceSet.Add(item)
		}
	}
	return symmetricDifferenceSet
}
