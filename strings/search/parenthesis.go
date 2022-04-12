package search

func Parenthesis(text string) bool {
	parcounter := int(0)

	for i := range text {
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
	return parcounter == 0
}
