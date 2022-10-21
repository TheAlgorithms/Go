package main

func cipher(text string, direction int) string {

	shift, offset := rune(3), rune(26)

	
	runes := []rune(text)

	for index, char := range runes {
		
		switch direction {
		case -1: 
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case +1: // decoding
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		
		runes[index] = char
	}

	return string(runes)
}
func encode(text string) string { return cipher(text, -1) }
func decode(text string) string { return cipher(text, +1) }

func main() {
	println("the text is `Hacktoberfest`")
	encoded := encode("das fuchedes 666")
	println("  encoded: " + encoded)
	decoded := decode(encoded)
	println("  decoded: " + decoded)
}
