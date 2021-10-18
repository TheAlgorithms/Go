package search

// Implementation of naive string search
// O(n*m) where n=len(txt) and m=len(pattern)
func Naive(text string, pattern string) []int {
	var positions []int
	for i := 0; i <= len(text)-len(pattern); i++ {
		var match bool = true
		for j := 0; j < len(pattern); j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}

		}
		if match {
			positions = append(positions, i)
		}
	}
	return positions
}
