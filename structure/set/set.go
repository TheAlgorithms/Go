// package set implements a Set using a golang map.
// This implies that only the types that are accepted as valid map keys can be used as set elements.
// For instance, do not try to Add a slice, or the program will panic.
package set

// New gives new set.
func New(items ...any) Set {
	st := set{
		elements: make(map[any]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

// Set is an interface of possible methods on 'set'.
type Set interface {
	// Add: adds new element to the set
	Add(item any)
	// Delete: deletes the passed element from the set if present
	Delete(item any)
	// Len: gives the length of the set (total no. of elements in set)
	Len() int
	// GetItems: gives the array( []any ) of elements of the set.
	GetItems() []any
	// In: checks whether item is present in set or not.
	In(item any) bool
	// IsSubsetOf: checks whether set is subset of set2 or not.
	IsSubsetOf(set2 Set) bool
	// IsProperSubsetOf: checks whether set is proper subset of set2 or not.
	// ex: [1,2,3] proper subset of [1,2,3,4] -> true
	IsProperSubsetOf(set2 Set) bool
	// IsSupersetOf: checks whether set is superset of set2 or not.
	IsSupersetOf(set2 Set) bool
	// IsProperSupersetOf: checks whether set is proper superset of set2 or not.
	// ex: [1,2,3,4] proper superset of [1,2,3] -> true
	IsProperSupersetOf(set2 Set) bool
	// Union: gives new union set of both sets.
	// ex: [1,2,3] union [3,4,5] -> [1,2,3,4,5]
	Union(set2 Set) Set
	// Intersection: gives new intersection set of both sets.
	// ex: [1,2,3] Intersection [3,4,5] -> [3]
	Intersection(set2 Set) Set
	// Difference: gives new difference set of both sets.
	// ex: [1,2,3] Difference [3,4,5] -> [1,2]
	Difference(set2 Set) Set
	// SymmetricDifference: gives new symmetric difference set of both sets.
	// ex: [1,2,3] SymmetricDifference [3,4,5] -> [1,2,4,5]
	SymmetricDifference(set2 Set) Set
}

type set struct {
	elements map[any]bool
}

func (st *set) Add(value any) {
	st.elements[value] = true
}

func (st *set) Delete(value any) {
	delete(st.elements, value)
}

func (st *set) GetItems() []any {
	keys := make([]any, 0, len(st.elements))
	for k := range st.elements {
		keys = append(keys, k)
	}
	return keys
}

func (st *set) Len() int {
	return len(st.elements)
}

func (st *set) In(value any) bool {
	if _, in := st.elements[value]; in {
		return true
	}
	return false
}

func (st *set) IsSubsetOf(superSet Set) bool {
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

func (st *set) IsProperSubsetOf(superSet Set) bool {
	if st.Len() == superSet.Len() {
		return false
	}
	return st.IsSubsetOf(superSet)
}

func (st *set) IsSupersetOf(subSet Set) bool {
	return subSet.IsSubsetOf(st)
}

func (st *set) IsProperSupersetOf(subSet Set) bool {
	if st.Len() == subSet.Len() {
		return false
	}
	return st.IsSupersetOf(subSet)
}

func (st *set) Union(st2 Set) Set {
	unionSet := New()
	for _, item := range st.GetItems() {
		unionSet.Add(item)
	}
	for _, item := range st2.GetItems() {
		unionSet.Add(item)
	}
	return unionSet
}

func (st *set) Intersection(st2 Set) Set {
	intersectionSet := New()
	var minSet, maxSet Set
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

func (st *set) Difference(st2 Set) Set {
	differenceSet := New()
	for _, item := range st.GetItems() {
		if !st2.In(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

func (st *set) SymmetricDifference(st2 Set) Set {
	symmetricDifferenceSet := New()
	dropSet := New()
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
