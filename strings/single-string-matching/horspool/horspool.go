package horspool

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// User defined.
// Set to true to read input from two command line arguments
// Set to false to read input from two files "pattern.txt" and "text.txt"
const isTakingInputFromCommandLine bool = false
const notFoundPosition int = -1

type result struct {
	foundPosition      int
	numberOfComparison int
}

// Implementation of Boyer-Moore-Horspool algorithm (Suffix based approach).
// Requires either a two command line arguments separated by a single space,
// or two files in the same folder: "pattern.txt" containing the string to
// be searched for, "text.txt" containing the text to be searched in.
func main() {
	var word string
	var text string

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
	fmt.Printf("\nRunning: Horspool algorithm.\n\n")
	fmt.Printf("Search word (%d chars long): %q.\n", len(word), word)
	fmt.Printf("Text        (%d chars long): %q.\n\n", len(text), text)
	r := horspool(text, word)
	if r.foundPosition == notFoundPosition {
		fmt.Printf("\n\nWord %q was not found.\n%d comparisons were done.", word, r.numberOfComparison)
	} else {
		fmt.Printf("\n\nWord %q was found at position %d in %q. \n%d comparisons were done.", word, r.foundPosition, text, r.numberOfComparison)
	}
}

// Function horspool performing the Horspool algorithm.
// Prints whether the word/pattern was found and on what position in the text or not.
func horspool(t, p string) result {
	m, n, c, pos := len(p), len(t), 0, 0
	//Preprocessing
	d := preprocess(t, p)
	//Map output
	fmt.Printf("Precomputed shifts per symbol: ")
	for key, value := range d {
		fmt.Printf("%c:%d; ", key, value)
	}
	fmt.Println()
	//Searching
	for pos <= n-m {
		j := m
		if t[pos+j-1] != p[j-1] {
			fmt.Printf("\n   comparing characters %c %c at positions %d %d", t[pos+j-1], p[j-1], pos+j-1, j-1)
			c++
		}
		for j > 0 && t[pos+j-1] == p[j-1] {
			fmt.Printf("\n   comparing characters %c %c at positions %d %d", t[pos+j-1], p[j-1], pos+j-1, j-1)
			c++
			fmt.Printf(" - match")
			j--
		}
		if j == 0 {
			return result{
				pos,
				c,
			}
		}
		pos = pos + d[t[pos+m]]
	}
	return result{
		notFoundPosition,
		c,
	}
}

// Function that pre-computes map with Key: uint8 (char) Value: int.
// Values determine safe shifting of search window.
func preprocess(t, p string) (d map[uint8]int) {
	d = make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		d[t[i]] = len(p)
	}
	for i := 0; i < len(p); i++ {
		d[p[i]] = len(p) - i
	}
	return d
}
