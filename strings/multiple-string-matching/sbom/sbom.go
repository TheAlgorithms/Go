package sbom

import (
	"fmt"
	"strings"
	"time"
)

// User defined.
// Set to true to print various extra stuff out (slows down the execution)
// Set to false for quick and quiet execution.
const debugMode bool = true

type Result struct {
	occurrences map[string][]int
}

// Implementation of Set Backward Oracle Matching algorithm (Factor based).
// Searches for a set of strings (in 'patterns.txt') in text.txt.
// func main() {
// 	patFile, err := ioutil.ReadFile("../patterns.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	textFile, err := ioutil.ReadFile("../text.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	patterns := strings.Split(string(patFile), " ")
// 	fmt.Printf("\nRunning: Set Backward Oracle Matching algorithm.\n\n")
// 	if debugMode {
// 		fmt.Printf("Searching for %d patterns/words:\n", len(patterns))
// 	}
// 	for i := 0; i < len(patterns); i++ {
// 		if len(patterns[i]) > len(textFile) {
// 			log.Fatal("There is a pattern that is longer than text! Pattern number:", i+1)
// 		}
// 		if debugMode {
// 			fmt.Printf("%q ", patterns[i])
// 		}
// 	}
// 	if debugMode {
// 		fmt.Printf("\n\nIn text (%d chars long): \n%q\n\n", len(textFile), textFile)
// 	}
// 	r := sbom(string(textFile), patterns)
// 	for key, value := range r.occurrences { //prints all occurrences of each pattern (if there was at least one)
// 		fmt.Printf("\nThere were %d occurences for word: %q at positions: ", len(value), key)
// 		for i := range value {
// 			fmt.Printf("%d", value[i])
// 			if i != len(value)-1 {
// 				fmt.Printf(", ")
// 			}
// 		}
// 		fmt.Printf(".")
// 	}

// }

// Sbom Function sbom performing the Set Backward Oracle Matching algorithm.
// Finds and prints occurrences of each pattern.
func Sbom(t string, p []string) Result {
	startTime := time.Now()
	occurrences := make(map[int][]int)
	lmin := computeMinLength(p)
	or, f := buildOracleMultiple(reverseAll(trimToLength(p, lmin)))
	if debugMode {
		fmt.Printf("\n\nSBOM:\n\n")
	}
	pos := 0
	for pos <= len(t)-lmin {
		current := 0
		j := lmin
		if debugMode {
			fmt.Printf("Position: %d, we read: ", pos)
		}
		for j >= 1 && stateExists(current, or) {
			if debugMode {
				fmt.Printf("%c", t[pos+j-1])
			}
			current = getTransition(current, t[pos+j-1], or)
			if debugMode {
				if current == -1 {
					fmt.Printf(" (FAIL) ")
				} else {
					fmt.Printf(", ")
				}
			}
			j--
		}
		if debugMode {
			fmt.Printf("in the factor oracle. \n")
		}
		word := getWord(pos, pos+lmin-1, t)
		if stateExists(current, or) && j == 0 && strings.HasPrefix(word, getCommonPrefix(p, f[current], lmin)) { //check for prefix match
			for i := range f[current] {
				if p[f[current][i]] == getWord(pos, pos-1+len(p[f[current][i]]), t) { //check for word match
					if debugMode {
						fmt.Printf("- Occurence, %q = %q\n", p[f[current][i]], word)
					}
					newOccurences := intArrayCapUp(occurrences[f[current][i]])
					occurrences[f[current][i]] = newOccurences
					occurrences[f[current][i]][len(newOccurences)-1] = pos
				}
			}
			j = 0
		}
		pos = pos + j + 1
	}
	elapsed := time.Since(startTime)
	fmt.Printf("\n\nElapsed %f secs\n", elapsed.Seconds())
	var resultOccurrences = make(map[string][]int)
	for key, value := range occurrences {
		resultOccurrences[p[key]] = value
	}

	return Result{
		resultOccurrences,
	}
}

// Function that builds factor oracle.
func buildOracleMultiple(p []string) (orToReturn map[int]map[uint8]int, f map[int][]int) {
	orTrie, stateIsTerminal, f := constructTrie(p)
	s := make([]int, len(stateIsTerminal)) //supply function
	i := 0                                 //root of trie
	orToReturn = orTrie
	s[i] = -1
	if debugMode {
		fmt.Printf("\n\nOracle construction: \n")
	}
	for current := 1; current < len(stateIsTerminal); current++ {
		o, parent := getParent(current, orTrie)
		down := s[parent]
		for stateExists(down, orToReturn) && getTransition(down, o, orToReturn) == -1 {
			createTransition(down, o, current, orToReturn)
			down = s[down]
		}
		if stateExists(down, orToReturn) {
			s[current] = getTransition(down, o, orToReturn)
		} else {
			s[current] = i
		}
	}
	return orToReturn, f
}

// Function that constructs Trie as an automaton for a set of reversed & trimmed strings.
func constructTrie(p []string) (trie map[int]map[uint8]int, stateIsTerminal []bool, f map[int][]int) {
	trie = make(map[int]map[uint8]int)
	stateIsTerminal = make([]bool, 1)
	f = make(map[int][]int)
	state := 1
	if debugMode {
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
			f[current] = newArray //F(Current) <- F(Current) union {i}
			if debugMode {
				fmt.Printf(" and %d", i)
			}
		} else {
			stateIsTerminal[current] = true
			f[current] = []int{i} //F(Current) <- {i}
			if debugMode {
				fmt.Printf("\n%d is terminal for word number %d", current, i)
			}
		}
	}
	return trie, stateIsTerminal, f
}

// Dynamically increases an array size of int's by 1.
func intArrayCapUp(old []int) (new []int) {
	new = make([]int, cap(old)+1)
	copy(new, old) //copy(dst,src)
	// old = new
	return new
}

// Dynamically increases an array size of bool's by 1.
func boolArrayCapUp(old []bool) (new []bool) {
	new = make([]bool, cap(old)+1)
	copy(new, old)
	// old = new
	return new
}

// Function that takes an array of strings and reverses it.
func reverseAll(s []string) (reversed []string) {
	reversed = make([]string, len(s))
	for i := 0; i < len(s); i++ {
		reversed[i] = reverse(s[i])
	}
	return reversed
}

// Function that takes a single string and reverses it.
func reverse(s string) string {
	l := len(s)
	m := make([]rune, l)
	for _, c := range s {
		l--
		m[l] = c
	}
	return string(m)
}

// Returns a prefix size 'lmin' for one string 'p' of first index found in 'f'.
// It is not needed to compare all the strings from 'p' indexed in 'f',
// thanks to the konwledge of 'lmin'.
func getCommonPrefix(p []string, f []int, lmin int) string {
	r := []rune(p[f[0]])
	newR := make([]rune, lmin)
	for j := 0; j < lmin; j++ {
		newR[j] = r[j]
	}
	return string(newR)
}

// Function that takes a set of strings 'p' and their wanted 'length'
// and then trims each string in that set to have desired 'length'.
func trimToLength(p []string, length int) (trimmedP []string) {
	trimmedP = make([]string, len(p))
	for i := range p {
		r := []rune(p[i])
		newR := make([]rune, length)
		for j := 0; j < length; j++ {
			newR[j] = r[j]
		}
		trimmedP[i] = string(newR)
	}
	return trimmedP
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

// Function that computes minimal length string in a set of strings.
func computeMinLength(p []string) (lmin int) {
	lmin = len(p[0])
	for i := 1; i < len(p); i++ {
		if len(p[i]) < lmin {
			lmin = len(p[i])
		}
	}
	return lmin
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
	if debugMode {
		fmt.Printf("\ncreated state %d", state)
	}
}

// Creates a transition for function σ(state,letter) = end.
func createTransition(fromState int, overChar uint8, toState int, at map[int]map[uint8]int) {
	at[fromState][overChar] = toState
	if debugMode {
		fmt.Printf("\n    σ(%d,%c)=%d;", fromState, overChar, toState)
	}
}

// Returns ending state for transition σ(fromState,overChar), '-1' if there is none.
func getTransition(fromState int, overChar uint8, at map[int]map[uint8]int) (toState int) {
	if !stateExists(fromState, at) {
		return -1
	}
	toState, ok := at[fromState][overChar]
	if ok {
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
