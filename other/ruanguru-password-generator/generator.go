package generator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func Reverse(str string) string {
	newStr := []string{}

	for i := len(str); i > 0; i-- {
		newStr = append(newStr, string(str[(i-1)]))
	}

	return strings.Join(newStr, "")
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Generate(str string) string {
	str = Reverse(str)
	newStr := ""
	// detect position of the vocal
	for i := 0; i < len(str); i++ {

		if string(str[i]) == "a" || string(str[i]) == "A" {
			if IsUpper(string(str[i])) {
				newStr += "e"
			} else {
				newStr += "E"
			}
			continue
		}
		if string(str[i]) == "e" || string(str[i]) == "E" {
			if IsUpper(string(str[i])) {
				newStr += "i"
			} else {
				newStr += "I"
			}
			continue
		}
		if string(str[i]) == "i" || string(str[i]) == "I" {
			if IsUpper(string(str[i])) {
				newStr += "o"
			} else {
				newStr += "O"
			}
			continue
		}
		if string(str[i]) == "o" || string(str[i]) == "O" {
			if IsUpper(string(str[i])) {
				newStr += "u"
			} else {
				newStr += "U"
			}

			continue
		}

		if string(str[i]) == "u" || string(str[i]) == "U" {
			if IsUpper(string(str[i])) {
				newStr += "a"
			} else {
				newStr += "A"
			}

			continue
		}

		if IsUpper(string(str[i])) {
			newStr += strings.ToLower(string(str[i]))
		} else {
			newStr += strings.ToUpper(string(str[i]))
		}
	}
	return strings.Join(strings.Split(newStr, " "), "")
}

func CheckPassword(str string) string {
	passwordLength := len(str)
	matchSymbol, _ := regexp.Compile(`[-~!@#$%^&*()_+{}|><?]+`)

	matchCapital, _ := regexp.Compile(`[a-z A-Z 0-9]+`)

	if passwordLength >= 14 && matchSymbol.MatchString(str) {
		return "strong"
	}

	if passwordLength > 6 && matchSymbol.MatchString(str) {
		return "average"
	}

	if passwordLength > 6 && matchCapital.MatchString(str) {
		return "weak"
	}

	if passwordLength < 7 {
		return "very weak"
	}

	return "error"

}

func PasswordGenerator(base string) (string, string) {
	password := Generate(base)

	securityLevel := CheckPassword(password)

	return password, securityLevel
}

func main() {
	// debug
	data := "YourDummyPassword"
	res, check := PasswordGenerator(data)

	fmt.Println(res)
	fmt.Println(check)
}
