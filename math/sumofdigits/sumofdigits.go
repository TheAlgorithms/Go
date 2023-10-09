// sumofdigits.go
// description: Finds the sum of digits of a whole number (base 10). If number is negetive returns -1.
// details:
// In mathematics, we can find the sum of digits of a whole number (integer >= 0). This is also known as Digit Sum.
// ref: (https://en.wikipedia.org/wiki/Digit_sum)
// time complexity: O(log10(N))
// space complexity: O(1)
// author: [SilverDragonOfR](https://github.com/SilverDragonOfR)

// Package sumofdigits provides method to find the sum of didits of a whole number
package sumofdigits

func Sumofdigits(number int) int {

	// handling the negetive number case
	if number < 0 {
		return -1
	}

	var sumOfDigits int = 0

	// to get the last digit till the number is not over and update number by shifting all digits to right
	for number > 0 {
		sumOfDigits += number % 10
		number /= 10
	}

	return sumOfDigits

}
