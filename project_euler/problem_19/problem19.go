package problem19

/**
* Problem 19 - Counting Sundays
* @see {@link https://projecteuler.net/problem=19}
*
* You are given the following information,
* but you may prefer to do some research for yourself.
*
* 1 Jan 1900 was a Monday.
* Thirty days has September,
* April, June and November.
* All the rest have thirty-one,
* Saving February alone,
* Which has twenty-eight, rain or shine.
* And on leap years, twenty-nine.
* A leap year occurs on any year evenly divisible by 4,
* but not on a century unless it is divisible by 400.
*
* How many Sundays fell on the first of the month during
* the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*
* @author ddaniel27
 */

func Problem19() int {
	count := 0
	dayOfWeek := 2 // 1 Jan 1901 was a Tuesday

	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			if dayOfWeek == 0 {
				count++
			}

			daysInMonth := 31
			switch month {
			case 4, 6, 9, 11:
				daysInMonth = 30
			case 2:
				if IsLeapYear(year) {
					daysInMonth = 29
				} else {
					daysInMonth = 28
				}
			}

			dayOfWeek = (dayOfWeek + daysInMonth) % 7
		}
	}

	return count
}

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}
