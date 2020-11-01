/*
This algorithm will convert an Eastern string number to an integer
https://en.wikipedia.org/wiki/Eastern_Arabic_numerals
Function receives a string as an Eastern number and outputs an integer
*/

package easterntointeger

var persianNumbers = map[string]int{"۰": 0, "۱": 1, "۲": 2, "۳": 3, "۴": 4, "۵": 5, "۶": 6, "۷": 7, "۸": 8, "۹": 9}
var arabicNumbers = map[string]int{"٠": 0, "١": 1, "٢": 2, "٣": 3, "٤": 4, "٥": 5, "٦": 6, "٧": 7, "٨": 8, "٩": 9}


func EasternToInteger(stringNumber string) int {
	if stringNumber == "" {
		return 0
	}
	number := 0
	for _, char := range stringNumber {
		if value, ok := persianNumbers[string(char)]; ok { // Persian check
			number = (number * 10) + value
		} else if value, ok := arabicNumbers[string(char)]; ok { // Arabic check
			number = (number * 10) + value
		}
	}
	return number
}
