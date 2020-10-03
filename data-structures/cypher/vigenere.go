package vigenere

func Sanitize(in string) string {
	out := []rune{}
	for _, v := range in {
		if 65 <= v && v <= 90 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}

	return string(out)
}

func Quartets(in string) string {
	out := make([]rune, 0, len(in))
	for i, v := range in {
		if i%4 == 0 && i != 0 {
			out = append(out, rune(32))
		}
		out = append(out, v)
	}
	return string(out)
}

func EncodePair(a, b rune) rune {
	return (((a - 'A') + (b - 'A')) % 26) + 'A'
}

func DecodePair(a, b rune) rune {
	return (((((a - 'A') - (b - 'A')) + 26) % 26) + 'A')
}

func Encipher(msg, key string) string {
	smsg, skey := Sanitize(msg), Sanitize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, EncodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func Decipher(msg, key string) string {
	smsg, skey := Sanitize(msg), Sanitize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, DecodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}
