package bom

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// User defined.
// Set to true to print various extra stuff out (slows down the execution)
// Set to false for quick and quiet execution.
const debugMode bool = false

type result struct {
	numberOfFoundWord int
	foundPositions    []int
}

// User defined.
// Set to true to read input from two command line arguments
// Set to false to read input from two files "pattern.txt" and "text.txt"
const isTakingInputFromCommandLine bool = false

// Implementation of Backward Oracle Matching algorithm (Factor based approach).
// Requires either a two command line arguments separated by a single space,
// or two files in the same folder: "pattern.txt" containing the string to
// be searched for, "text.txt" containing the text to be searched in.
func main() {
	var word string
	var text string
	//occurrences

	if isTakingInputFromCommandLine { // case of command line input
		args := os.Args
		if len(args) <= 2 {
			log.Fatal("Not enough arguments. Two string arguments separated by spaces are required!")
		}
		word = args[1]
		text = args[2]
		for i := 3; i < len(args); i++ {
			text = text + " " + args[i]
		}
	} else { // case of file line input
		patFile, err := ioutil.ReadFile("../pattern.txt")
		if err != nil {
			log.Fatal(err)
		}
		textFile, err := ioutil.ReadFile("../text.txt")
		if err != nil {
			log.Fatal(err)
		}
		word = string(patFile)
		text = string(textFile)
	}

	if len(word) > len(text) {
		log.Fatal("Pattern is longer than text!")
	}
	if debugMode {
		fmt.Printf("\nRunning: Backward Oracle Matching alghoritm.\n\n")
		fmt.Printf("Search word (%d chars long): %q.\n", len(word), word)
		fmt.Printf("Text        (%d chars long): %q.\n\n", len(text), text)
	} else {
		fmt.Printf("\nRunning: Backward Oracle Matching alghoritm.\n\n")
	}
	r := bom(text, word)
	if r.numberOfFoundWord == 0 {
		fmt.Printf("\nWord was not found.\n")
	} else {
		fmt.Printf("Word %q was found %d times at positions: ", word, r.numberOfFoundWord)
		for k := 0; k < r.numberOfFoundWord; k++ {
			fmt.Printf("%d", r.foundPositions[k])
			if k < r.numberOfFoundWord-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf(".\n")
	}
}

// Function bom performing the Backward Oracle Matching algorithm.
// Prints whether the word/pattern was found + positions of possible multiple occurrences
// or that the word was not found.
func bom(t, p string) result {
	startTime := time.Now()
	n, m := len(t), len(p)
	var current, j, pos int
	oracle := oracleOnLine(reverse(p))
	occurrences := make([]int, len(t))
	currentOcc := 0
	pos = 0
	if debugMode == true {
		fmt.Printf("\n\nWe are reading backwards in %q, searching for %q\n\nat position %d:\n", t, p, pos+m-1)
	}
	for pos <= n-m {
		current = 0 //initial state of the oracle
		j = m
		for j > 0 && stateExists(current, oracle) {
			if debugMode == true {
				prettyPrint(current, j, n, pos, t, oracle)
			}
			current = getTransition(current, t[pos+j-1], oracle)
			j--
		}
		if stateExists(current, oracle) {
			if debugMode == true {
				fmt.Printf(" We got an occurence!")
			}
			occurrences[currentOcc] = pos
			currentOcc++
		}
		pos = pos + j + 1
		if pos+m-1 < len(t) {
			if debugMode == true {
				fmt.Printf("\n\nposition %d:\n", pos+m-1)
			}
		}
	}
	elapsed := time.Since(startTime)
	fmt.Printf("\n\nElapsed %f secs\n", elapsed.Seconds())
	fmt.Printf("\n\n")
	if currentOcc > 0 {
		results := make([]int, currentOcc)
		for i := 0; i < currentOcc; i++ {
			results[i] = occurrences[i]
		}
		return result{currentOcc, results}
	}
	return result{0, []int{}}
}

// Construction of the factor oracle automaton for a word p.
func oracleOnLine(p string) (oracle map[int]map[uint8]int) {
	if debugMode == true {
		fmt.Printf("Oracle construction: \n")
	}
	oracle = make(map[int]map[uint8]int)
	supply := make([]int, len(p)+2) // supply function
	createNewState(0, oracle)
	supply[0] = -1
	var orP string
	for j := 0; j < len(p); j++ {
		oracle, orP = oracleAddLetter(oracle, supply, orP, p[j])
	}
	return oracle
}

// Adds one letter to the oracle.
func oracleAddLetter(oracle map[int]map[uint8]int, supply []int, orP string, o uint8) (oracleToReturn map[int]map[uint8]int, orPToReturn string) {
	m := len(orP)
	var s int
	createNewState(m+1, oracle)
	createTransition(m, o, m+1, oracle)
	k := supply[m]
	for k > -1 && getTransition(k, o, oracle) == -1 {
		createTransition(k, o, m+1, oracle)
		k = supply[k]
	}
	if k == -1 {
		s = 0
	} else {
		s = getTransition(k, o, oracle)
	}
	supply[m+1] = s
	return oracle, orP + string(o)
}

// Function that takes a single string and reverses it.
// @author 'Walter' http://stackoverflow.com/a/10043083
func reverse(s string) string {
	l := len(s)
	m := make([]rune, l)
	for _, c := range s {
		l--
		m[l] = c
	}
	return string(m)
}

// Automaton function for creating a new state.
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

// Returns ending state for transition σ(fromState,overChar), -1 if there is none.
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

// Checks if state exists. Returns true if it does, false otherwise.
func stateExists(state int, at map[int]map[uint8]int) bool {
	_, ok := at[state]
	if !ok || state == -1 || at[state] == nil {
		return false
	}
	return true
}

// Just some printing of extra information about what the algorithm does.
func prettyPrint(current int, j int, n int, pos int, t string, oracle map[int]map[uint8]int) {
	if current == 0 && !(getTransition(current, t[pos+j-1], oracle) == -1) {
		fmt.Printf("\n -->(%d)---(%c)--->(%d)", current, t[pos+j-1], getTransition(current, t[pos+j-1], oracle))
	} else if getTransition(current, t[pos+j-1], oracle) == -1 && current != 0 {
		fmt.Printf("\n    (%d)---(%c)       ", current, t[pos+j-1])
	} else if getTransition(current, t[pos+j-1], oracle) == -1 && current == 0 {
		fmt.Printf("\n -->(%d)---(%c)       ", current, t[pos+j-1])
	} else {
		fmt.Printf("\n    (%d)---(%c)--->(%d)", current, t[pos+j-1], getTransition(current, t[pos+j-1], oracle))
	}
	fmt.Printf(" ")
	for a := 0; a < pos+j-1; a++ {
		fmt.Printf("%c", t[a])
	}
	if getTransition(current, t[pos+j-1], oracle) == -1 {
		fmt.Printf("[%c]", t[pos+j-1])
	} else {
		fmt.Printf("[%c]", t[pos+j-1])
	}
	for a := pos + j; a < n; a++ {
		fmt.Printf("%c", t[a])
	}
	if getTransition(current, t[pos+j-1], oracle) == -1 {
		fmt.Printf(" FAIL on the character[%c]", t[pos+j-1])
	}
}
