package set

// GetSet gives new set.
func GetSet(items ...interface{}) Set {
	st := set{
		elements: make(map[interface{}]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

// Set is an interface of possible methods on 'set'.
type Set interface {
	// Add: adds new element to the set
	Add(item interface{})
	// Delete: delets the passed element from the set if present
	Delete(item interface{})
	// Len: gives the length of the set (total no. of elements in set)
	Len() int
	// GetItems: gives the array( []interface{} ) of elements of the set.
	GetItems() []interface{}
	// In: checks whether item is present in set or not.
	In(item interface{}) bool
	// IsSubsetOf: checks wether set is subset of set2 or not.
	IsSubsetOf(set2 Set) bool
	// IsSupersetOf: checks wether set is superset of set2 or not.
	IsSupersetOf(set2 Set) bool
	// Union: gives new union set of both sets.
	// ex: [1,2,3] union [3,4,5] -> [1,2,3,4,5]
	Union(set2 Set) Set
	// Intersection: gives new intersection set of both sets.
	// ex: [1,2,3] union [3,4,5] -> [3]
	Intersection(set2 Set) Set
	// Difference: gives new difference set of both sets.
	// ex: [1,2,3] union [3,4,5] -> [1,2]
	Difference(set2 Set) Set
	// SymmetricDifference: gives new symmetric difference set of both sets.
	// ex: [1,2,3] union [3,4,5] -> [1,2,4,5]
	SymmetricDifference(set2 Set) Set
}

type set struct {
	elements map[interface{}]bool
}

func (st *set) Add(value interface{}) {
	st.elements[value] = true
}

func (st *set) Delete(value interface{}) {
	delete(st.elements, value)
}

func (st *set) GetItems() []interface{} {
	keys := make([]interface{}, 0, len(st.elements))
	for k := range st.elements {
		keys = append(keys, k)
	}
	return keys
}

func (st *set) Len() int {
	return len(st.elements)
}

func (st *set) In(value interface{}) bool {
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

func (st *set) IsSupersetOf(subSet Set) bool {
	return subSet.IsSubsetOf(st)
}

func (st *set) Union(st2 Set) Set {
	unionSet := GetSet()
	for _, item := range st.GetItems() {
		unionSet.Add(item)
	}
	for _, item := range st2.GetItems() {
		unionSet.Add(item)
	}
	return unionSet
}

func (st *set) Intersection(st2 Set) Set {
	intersectionSet := GetSet()
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
	differenceSet := GetSet()
	for _, item := range st.GetItems() {
		if !st2.In(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

func (st *set) SymmetricDifference(st2 Set) Set {
	symmetricDifferenceSet := GetSet()
	dropSet := GetSet()
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
