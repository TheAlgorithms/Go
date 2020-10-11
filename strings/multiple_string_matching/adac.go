package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// User defined.
// Set to true to print various extra stuff out (slows down the execution)
// Set to false for quick and quiet execution.
const debugMode bool = true

// Implementation of Advanced Aho-Corasick algorithm (Prefix based).
// Searches for a set of strings (patterns.txt) in text.txt.
func main() {
	patFile, err := ioutil.ReadFile("patterns.txt")
	if err != nil {
		log.Fatal(err)
	}
	textFile, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	patterns := strings.Split(string(patFile), " ")
	fmt.Printf("\nRunning: Advanced Aho-Corasick algorithm.\n\n")
	if debugMode == true {
		fmt.Printf("Searching for %d patterns/words:\n", len(patterns))
	}
	for i := 0; i < len(patterns); i++ {
		if len(patterns[i]) > len(textFile) {
			log.Fatal("There is a pattern that is longer than text! Pattern number:", i+1)
		}
		if debugMode == true {
			fmt.Printf("%q ", patterns[i])
		}
	}
	if debugMode == true {
		fmt.Printf("\n\nIn text (%d chars long): \n%q\n\n", len(textFile), textFile)
	}
	ahoCorasick(string(textFile), patterns)
}

// Function performing the Advanced Aho-Corasick alghoritm.
// Finds and prints occurences of each pattern.
func ahoCorasick(t string, p []string) {
	startTime := time.Now()
	occurences := make(map[int][]int)
	ac, f := buildExtendedAc(p)
	if debugMode == true {
		fmt.Printf("\n\nAC:\n\n")
	}
	current := 0
	for pos := 0; pos < len(t); pos++ {
		if getTransition(current, t[pos], ac) != -1 {
			current = getTransition(current, t[pos], ac)
		} else {
			current = 0
		}
		_, ok := f[current]
		if ok {
			for i := range f[current] {
				if p[f[current][i]] == getWord(pos-len(p[f[current][i]])+1, pos, t) { //check for word match
					if debugMode == true {
						fmt.Printf("Occurence at position %d, %q = %q\n", pos-len(p[f[current][i]])+1, p[f[current][i]], p[f[current][i]])
					}
					newOccurences := intArrayCapUp(occurences[f[current][i]])
					occurences[f[current][i]] = newOccurences
					occurences[f[current][i]][len(newOccurences)-1] = pos - len(p[f[current][i]]) + 1
				}
			}
		}
	}
	elapsed := time.Since(startTime)
	fmt.Printf("\n\nElapsed %f secs\n", elapsed.Seconds())
	for key, value := range occurences { //prints all occurences of each pattern (if there was at least one)
		fmt.Printf("\nThere were %d occurences for word: %q at positions: ", len(value), p[key])
		for i := range value {
			fmt.Printf("%d", value[i])
			if i != len(value)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf(".")
	}
	return
}

// Functions that builds extended Aho Corasick automaton.
func buildExtendedAc(p []string) (acToReturn map[int]map[uint8]int, f map[int][]int) {
	acTrie, stateIsTerminal, f := constructTrie(p)
	s := make([]int, len(stateIsTerminal)) //supply function
	i := 0                                 //root of acTrie
	acToReturn = acTrie
	s[i] = -1
	if debugMode == true {
		fmt.Printf("\n\nAC construction: \n")
	}
	for current := 1; current < len(stateIsTerminal); current++ {
		o, parent := getParent(current, acTrie)
		down := s[parent]
		for stateExists(down, acToReturn) && getTransition(down, o, acToReturn) == -1 {
			down = s[down]
		}
		if stateExists(down, acToReturn) {
			s[current] = getTransition(down, o, acToReturn)
			if stateIsTerminal[s[current]] == true {
				stateIsTerminal[current] = true
				f[current] = arrayUnion(f[current], f[s[current]]) //F(Current) <- F(Current) union F(S(Current))
				if debugMode == true {
					fmt.Printf(" f[%d] set to: ", current)
					for i := range f[current] {
						fmt.Printf("%d\n", f[current][i])
					}
				}
			}
		} else {
			s[current] = i //initial state?
		}
	}
	if debugMode == true {
		fmt.Printf("\nsupply function: \n")
		for i := range s {
			fmt.Printf("\ns[%d]=%d", i, s[i])
		}
		fmt.Printf("\n\n")
		for i, j := range f {
			fmt.Printf("f[%d]=", i)
			for k := range j {
				fmt.Printf("%d\n", j[k])
			}
		}
	}
	if debugMode == true {
		fmt.Printf("\n\nAdAC completion: \n")
	}
	// advanced Aho-Corasick part
	a := computeAlphabet(p) // concat of all patterns in p
	for j := range a {
		if getTransition(i, a[j], acToReturn) == -1 {
			createTransition(i, a[j], i, acToReturn)
		}
	}
	for current := 1; current < len(stateIsTerminal); current++ {
		for j := range a {
			if getTransition(current, a[j], acToReturn) == -1 {
				createTransition(current, a[j], getTransition(s[current], a[j], acToReturn), acToReturn)
			}
		}
	}
	return acToReturn, f
}

// Function that constructs Trie as an automaton for a set of reversed & trimmed strings.
func constructTrie(p []string) (trie map[int]map[uint8]int, stateIsTerminal []bool, f map[int][]int) {
	trie = make(map[int]map[uint8]int)
	stateIsTerminal = make([]bool, 1)
	f = make(map[int][]int)
	state := 1
	if debugMode == true {
		fmt.Printf("\n\nTrie construction: \n")
	}
	createNewState(0, trie)
	for i := 0; i < len(p); i++ {
		current := 0
		j := 0
		for j < len(p[i]) && getTransition(current, p[i][j], trie) != -1 {
			current = getTransition(current, p[i][j], trie)
			j++
		}
		for j < len(p[i]) {
			stateIsTerminal = boolArrayCapUp(stateIsTerminal)
			createNewState(state, trie)
			stateIsTerminal[state] = false
			createTransition(current, p[i][j], state, trie)
			current = state
			j++
			state++
		}
		if stateIsTerminal[current] {
			newArray := intArrayCapUp(f[current])
			newArray[len(newArray)-1] = i
			f[current] = newArray // F(Current) <- F(Current) union {i}
			if debugMode == true {
				fmt.Printf(" and %d", i)
			}
		} else {
			stateIsTerminal[current] = true
			f[current] = []int{i} // F(Current) <- {i}
			if debugMode == true {
				fmt.Printf("\n%d is terminal for word number %d", current, i)
			}
		}
	}
	return trie, stateIsTerminal, f
}

// Returns 'true' if arry of int's 's' contains int 'e', 'false' otherwise.
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Function that returns word found in text 't' at position range 'begin' to 'end'.
func getWord(begin, end int, t string) string {
	for end >= len(t) {
		return ""
	}
	d := make([]uint8, end-begin+1)
	for j, i := 0, begin; i <= end; i, j = i+1, j+1 {
		d[j] = t[i]
	}
	return string(d)
}

// Function that returns string of all the possible characters in given patterns.
func computeAlphabet(p []string) (s string) {
	s = p[0]
	for i := 1; i < len(p); i++ {
		s = s + p[i]
	}
	return s
}

// Dynamically increases an array size of int's by 1.
func intArrayCapUp(old []int) (new []int) {
	new = make([]int, cap(old)+1)
	copy(new, old) //copy(dst,src)
	old = new
	return new
}

// Dynamically increases an array size of bool's by 1.
func boolArrayCapUp(old []bool) (new []bool) {
	new = make([]bool, cap(old)+1)
	copy(new, old)
	old = new
	return new
}

// Concats two arrays of int's into one.
func arrayUnion(to, from []int) (concat []int) {
	concat = to
	for i := range from {
		if !contains(concat, from[i]) {
			concat = intArrayCapUp(concat)
			concat[len(concat)-1] = from[i]
		}
	}
	return concat
}

// Function that finds the first previous state of a state and returns it.
// Used for trie where there is only one parent.
func getParent(state int, at map[int]map[uint8]int) (uint8, int) {
	for beginState, transitions := range at {
		for c, endState := range transitions {
			if endState == state {
				return c, beginState
			}
		}
	}
	return 0, 0 //unreachable
}

// Automaton function for creating a new state 'state'.
func createNewState(state int, at map[int]map[uint8]int) {
	at[state] = make(map[uint8]int)
	if debugMode == true {
		fmt.Printf("\ncreated state %d", state)
	}
}

// Creates a transition for function σ(state,letter) = end.
func createTransition(fromState int, overChar uint8, toState int, at map[int]map[uint8]int) {
	at[fromState][overChar] = toState
	if debugMode == true {
		fmt.Printf("\n    σ(%d,%c)=%d;", fromState, overChar, toState)
	}
}

// Returns ending state for transition σ(fromState,overChar), '-1' if there is none.
func getTransition(fromState int, overChar uint8, at map[int]map[uint8]int) (toState int) {
	if !stateExists(fromState, at) {
		return -1
	}
	toState, ok := at[fromState][overChar]
	if ok == false {
		return -1
	}
	return toState
}

// Checks if state 'state' exists. Returns 'true' if it does, 'false' otherwise.
func stateExists(state int, at map[int]map[uint8]int) bool {
	_, ok := at[state]
	if !ok || state == -1 || at[state] == nil {
		return false
	}
	return true
}
