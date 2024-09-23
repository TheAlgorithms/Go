package dynamic

func WordBreak(s string, wordDict []string) bool {
	return helper(s, wordDict, 0, map[int]bool{})
}

func helper(s string, wordDict []string, i int, dp map[int]bool) bool {
	if val, ok := dp[i]; ok {
		return val
	}
	if i == len(s) {
		dp[i] = true
		return dp[i]
	}
	for _, word := range wordDict {
		if len(s[i:]) < len(word) {
			continue
		}
		if s[i:i+len(word)] != word {
			continue
		}
		if v := helper(s, wordDict, i+len(word), dp); v {
			dp[i] = v
			return true
		}
	}
	dp[i] = false
	return dp[i]
}
