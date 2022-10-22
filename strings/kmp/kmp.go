package kmp

import (
	"fmt"
)

const notFoundPosition int = -1

type Result struct {
	resultPosition     int
	numberOfComparison int
}

// Kmp Function kmp performing the Knuth-Morris-Pratt algorithm.
// Prints whether the word/pattern was found and on what position in the text or not.
// m - current match in text, i - current character in w, c - amount of comparisons.
func Kmp(text string, word string) Result {
	m, i, c := 0, 0, 0
	t := kmpTable(word)
	for m+i < len(text) {
		fmt.Printf("\n   comparing characters %c %c at positions %d %d", text[m+i], word[i], m+i, i)
		c++
		if word[i] == text[m+i] {
			fmt.Printf(" - match")
			if i == len(word)-1 {
				return Result{
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
	return Result{notFoundPosition,
		c,
	}
}

// Table building algorithm.
// Takes word to be analyzed and table to be filled.
func kmpTable(word string) (t []int) {
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
