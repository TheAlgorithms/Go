package ahocorasick

// ConstructTrie Function that constructs Trie as an automaton for a set of reversed & trimmed strings.
func ConstructTrie(p []string) (trie map[int]map[uint8]int, stateIsTerminal []bool, f map[int][]int) {
	trie = make(map[int]map[uint8]int)
	stateIsTerminal = make([]bool, 1)
	f = make(map[int][]int)
	state := 1
	CreateNewState(0, trie)
	for i := 0; i < len(p); i++ {
		current := 0
		j := 0
		for j < len(p[i]) && GetTransition(current, p[i][j], trie) != -1 {
			current = GetTransition(current, p[i][j], trie)
			j++
		}
		for j < len(p[i]) {
			stateIsTerminal = BoolArrayCapUp(stateIsTerminal)
			CreateNewState(state, trie)
			stateIsTerminal[state] = false
			CreateTransition(current, p[i][j], state, trie)
			current = state
			j++
			state++
		}
		if stateIsTerminal[current] {
			newArray := IntArrayCapUp(f[current])
			newArray[len(newArray)-1] = i
			f[current] = newArray // F(Current) <- F(Current) union {i}
		} else {
			stateIsTerminal[current] = true
			f[current] = []int{i} // F(Current) <- {i}
		}
	}
	return trie, stateIsTerminal, f
}

// Contains Returns 'true' if array of int's 's' contains int 'e', 'false' otherwise.
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GetWord Function that returns word found in text 't' at position range 'begin' to 'end'.
func GetWord(begin, end int, t string) string {
	for end >= len(t) {
		return ""
	}
	d := make([]uint8, end-begin+1)
	for j, i := 0, begin; i <= end; i, j = i+1, j+1 {
		d[j] = t[i]
	}
	return string(d)
}

// ComputeAlphabet Function that returns string of all the possible characters in given patterns.
func ComputeAlphabet(p []string) (s string) {
	s = p[0]
	for i := 1; i < len(p); i++ {
		s = s + p[i]
	}
	return s
}

// IntArrayCapUp Dynamically increases an array size of int's by 1.
func IntArrayCapUp(old []int) (new []int) {
	new = make([]int, cap(old)+1)
	copy(new, old) //copy(dst,src)
	// old = new
	return new
}

// BoolArrayCapUp Dynamically increases an array size of bool's by 1.
func BoolArrayCapUp(old []bool) (new []bool) {
	new = make([]bool, cap(old)+1)
	copy(new, old)
	// old = new
	return new
}

// ArrayUnion Concats two arrays of int's into one.
func ArrayUnion(to, from []int) (concat []int) {
	concat = to
	for i := range from {
		if !Contains(concat, from[i]) {
			concat = IntArrayCapUp(concat)
			concat[len(concat)-1] = from[i]
		}
	}
	return concat
}

// GetParent Function that finds the first previous state of a state and returns it.
// Used for trie where there is only one parent.
func GetParent(state int, at map[int]map[uint8]int) (uint8, int) {
	for beginState, transitions := range at {
		for c, endState := range transitions {
			if endState == state {
				return c, beginState
			}
		}
	}
	return 0, 0 //unreachable
}

// CreateNewState Automaton function for creating a new state 'state'.
func CreateNewState(state int, at map[int]map[uint8]int) {
	at[state] = make(map[uint8]int)
}

// CreateTransition Creates a transition for function σ(state,letter) = end.
func CreateTransition(fromState int, overChar uint8, toState int, at map[int]map[uint8]int) {
	at[fromState][overChar] = toState
}

// GetTransition Returns ending state for transition σ(fromState,overChar), '-1' if there is none.
func GetTransition(fromState int, overChar uint8, at map[int]map[uint8]int) (toState int) {
	if !StateExists(fromState, at) {
		return -1
	}
	toState, ok := at[fromState][overChar]
	if !ok {
		return -1
	}
	return toState
}

// StateExists Checks if state 'state' exists. Returns 'true' if it does, 'false' otherwise.
func StateExists(state int, at map[int]map[uint8]int) bool {
	_, ok := at[state]
	if !ok || state == -1 || at[state] == nil {
		return false
	}
	return true
}
