package kmp

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// User defined.
// Set to true to read input from two command line arguments
// Set to false to read input from two files "pattern.txt" and "text.txt"
const isTakingInputFromCommandLine bool = true
const notFoundPosition int = -1

type result struct {
	resultPosition     int
	numberOfComparison int
}

// Implementation of Knuth-Morris-Pratt algorithm (Prefix based approach).
// Requires either a two command line arguments separated by a single space,
// or two files in the same folder: "pattern.txt" containing the string to
// be searched for, "text.txt" containing the text to be searched in.
func main() {
	var text string
	var word string

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
	} else { // case of file input
		patFile, err := ioutil.ReadFile("../pattern.txt")
		if err != nil {
			log.Fatal(err)
		}
		textFile, err := ioutil.ReadFile("../text.txt")
		if err != nil {
			log.Fatal(err)
		}
		text = string(textFile)
		word = string(patFile)
	}

	if len(word) > len(text) {
		log.Fatal("Pattern is longer than text!")
	}
	fmt.Printf("\nRunning: Knuth-Morris-Pratt algorithm.\n\n")
	fmt.Printf("Search word (%d chars long): %q.\n", len(word), word)
	fmt.Printf("Text        (%d chars long): %q.\n\n", len(text), text)

	r := kmp(text, word)
	if r.resultPosition == notFoundPosition {
		fmt.Printf("\n\nWord was not found.\n%d comparisons were done.", r.numberOfComparison)
	} else {
		fmt.Printf("\n\nWord %q was found at position %d in %q. \n%d comparisons were done.", word,
			r.resultPosition, text, r.numberOfComparison)
	}
}

// Function kmp performing the Knuth-Morris-Pratt algorithm.
// Prints whether the word/pattern was found and on what position in the text or not.
// m - current match in text, i - current character in w, c - amount of comparisons.
func kmp(text string, word string) result {
	m, i, c := 0, 0, 0
	t := kmp_table(word)
	for m+i < len(text) {
		fmt.Printf("\n   comparing characters %c %c at positions %d %d", text[m+i], word[i], m+i, i)
		c++
		if word[i] == text[m+i] {
			fmt.Printf(" - match")
			if i == len(word)-1 {
				return result{
					m, c,
				}
			}
			i++
		} else {
			m = m + i - t[i]
			if t[i] > -1 {
				i = t[i]
			} else {
				i = 0
			}
		}
	}
	return result{notFoundPosition,
		c,
	}
}

// Table building alghoritm.
// Takes word to be analyzed and table to be filled.
func kmp_table(word string) (t []int) {
	t = make([]int, len(word))
	pos, cnd := 2, 0
	t[0], t[1] = -1, 0
	for pos < len(word) {
		if word[pos-1] == word[cnd] {
			cnd++
			t[pos] = cnd
			pos++
		} else if cnd > 0 {
			cnd = t[cnd]
		} else {
			t[pos] = 0
			pos++
		}
	}
	return t
}
