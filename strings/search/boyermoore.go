package search

// Implementation of boyer moore string search
// O(l) where l=len(text)
func BoyerMoore(text string, pattern string) []int {
	var positions []int

	l := len(text)
	n := len(pattern)

	// using booyer moore horspool modification
	// O(n) space instead of O(n**2)
	bcr := make(map[byte]int)
	for i := 0; i < n-1; i++ {
		bcr[pattern[i]] = n - i - 1
	}

	// Apostolico–Giancarlo modification
	// allow to skip patterns that we know matches
	// let us do O(l) instead of O(ln)
	skips := make(map[int]int)
	for _, s := range bcr {
		i := 0
		for ; i < n-s; i++ {
			if pattern[n-1-i] != pattern[n-1-s-i] {
				break
			}
		}
		skips[s] = i
	}

	skip := 0
	jump := n
	for i := 0; i < l-n+1; {
		skip = skips[jump]
		for k := n - 1; k > -1; k-- {
			if text[i+k] != pattern[k] {
				jump, ok := bcr[text[i+k]]
				if !ok {
					jump = n
				}
				i += jump
				break
			}
			if k == n-jump {
				k -= skip
			}
			if k == 0 {
				positions = append(positions, i)
				jump = 1
				i += jump
			}
		}
	}

	return positions
}
