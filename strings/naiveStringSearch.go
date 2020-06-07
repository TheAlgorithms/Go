/*
This algorithm tries to search the given pattern in the given text.
If pattern is found from index position i in the text, 
it is added to positions.

Time Complexity : O(n*m)
    n = length of text
    m = length of pattern 
*/

package main

import (  
    "fmt"
)

func naivePatternSearch(text string, pattern string) []int { 
    var positions []int 
    for i := 0; i < len(text)-len(pattern); i++ {
        var match = true
        j := i +(len(pattern))
        if text[i:j] != pattern {
                match = false
                continue
            }        
        if match {
            positions = append(positions, i)
        }
    }   
    return positions
}

func main() {  
    text := "ABAAABCDBBABCDDEBCABC"
    pattern := "ABC"
    positions := naivePatternSearch(text, pattern)
    if len(positions) == 0 {
        fmt.Println("Pattern not found in given text!")
    } else {
        fmt.Println("Pattern found in following position(s):")
        fmt.Printf("%v \n", positions)
        fmt.Println("Note that position is zero indexed :)")
    }
}
