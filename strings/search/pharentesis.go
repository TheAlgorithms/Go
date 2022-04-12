package search

func Pharentesis(text string) bool {
	var phar_counter int
	phar_counter = 0

	for i := 0; i < len(text); i++ {
		switch text[i] {
			case '(':
				phar_counter++
			case ')':
				phar_counter--
		}
		if phar_counter < 0 {
			return false
		}
	}
	if phar_counter == 0{
		return true
	}
	return false
}
