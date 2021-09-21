// spigotpi.go
// description: A Spigot Algorithm for the Digits of Pi
// details:
// implementation of Spigot Algorithm for the Digits of Pi - [Spigot algorithm](https://en.wikipedia.org/wiki/Spigot_algorithm)
// author(s) [red_byte](https://github.com/i-redbyte)
// see spigotpi_test.go

package pi

import "strconv"

func PiSpigot(n int) string {
	pi := ""
	boxes := n * 10 / 3
	reminders := make([]int, boxes)
	for i := 0; i < boxes; i++ {
		reminders[i] = 2
	}
	heldDigits := 0
	for i := 0; i < n; i++ {
		carriedOver := 0
		sum := 0
		for j := boxes - 1; j >= 0; j-- {
			reminders[j] *= 10
			sum = reminders[j] + carriedOver
			quotient := sum / (j*2 + 1)
			reminders[j] = sum % (j*2 + 1)
			carriedOver = quotient * j
		}
		reminders[0] = sum % 10
		q := sum / 10
		if q == 9 {
			heldDigits++
		} else if q == 10 {
			q = 0
			for k := 1; k <= heldDigits; k++ {
				replaced, _ := strconv.Atoi(pi[i-k : i-k+1])
				if replaced == 9 {
					replaced = 0
				} else {
					replaced++
				}
				pi = string(delChar([]rune(pi), i-k))
				pi = pi[:i-k] + strconv.Itoa(replaced) + pi[i-k:]
			}
			heldDigits = 1
		} else {
			heldDigits = 1
		}
		pi += strconv.Itoa(q)
	}
	return pi
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}
