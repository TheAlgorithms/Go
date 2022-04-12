package search

func Parenthesis(text string) bool {
	parcounter := int(0)

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '(':
			parcounter++
		case ')':
			parcounter--
		}
		if parcounter < 0 {
			return false
		}
	}
	if parcounter == 0 {
		return true
	}
	return false
}
