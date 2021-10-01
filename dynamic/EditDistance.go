/**
 * A DynamicProgramming based solution for Edit Distance problem In Go Description of Edit
 * Distance with an Example:
 *
 * <p>Edit distance is a way of quantifying how dissimilar two strings (e.g., words) are to one
 * another, by counting the minimum number of operations required to transform one string into the
 * other. The distance operations are the removal, insertion, or substitution of a character in the
 * string.
 *
 * <p>
 *
 * <p>The Distance between "kitten" and "sitting" is 3. A minimal edit script that transforms the
 * former into the latter is:
 *
 * <p>kitten → sitten (substitution of "s" for "k") sitten → sittin (substitution of "i" for "e")
 * sittin → sitting (insertion of "g" at the end).
 *
 */
package main
import "fmt"
func minOf(vars ...int) int{ // this function gives the minimum value
    min := vars[0]
    
    for _, val := range vars{
        if val < min{
            min = val     
        }
    }
    
    return min
}
func minDistance(word1 string, word2 string) int {
    len1, len2 := len(word1), len(word2) 
    
    costs := make([][]int, len1 + 1)
    for idx := range costs{
        costs[idx] = make([]int, len2 + 1) //we make a 2d array of size (len1+1)x(len2+1)
    }
    
    for idx := 0; idx <= len1; idx++ {  
        costs[idx][0] = idx
    }
    /*If first string is empty, the only option is to 
	insert all characters of second string into first */
    for idx := 0; idx <= len2; idx++{
        costs[0][idx] = idx
    }
    // iterate though, and check last char
    for idx1 := 0; idx1 < len1; idx1++{
        for idx2 := 0; idx2 < len2; idx2++ {
            if word1[idx1] == word2[idx2] {
                costs[idx1 + 1][idx2 + 1] = costs[idx1][idx2] // // update costs value for +1 length
            }else{
				/* if two characters are different ,
          then take the minimum of the various operations(i.e insertion,removal,substitution)*/
                replacement := costs[idx1][idx2]
                insersion := costs[idx1 + 1][idx2]
                deletion := costs[idx1][idx2 + 1]
                
                costs[idx1 + 1][idx2 + 1] = 1 + minOf(replacement, insersion, deletion)
            }
        }
    }
    
    return costs[len1][len2]
	 /* return the final answer , after traversing through both the strings*/
}
/*
func main(){
    word1:="horse"
    word2:="ros"
    fmt.Print(minDistance(word1, word2), "\n")
}
*/