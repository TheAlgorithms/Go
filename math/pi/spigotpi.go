// spigotpi.go
// description: A Spigot Algorithm for the Digits of Pi
// details:
// implementation of Spigot Algorithm for the Digits of Pi - [Spigot algorithm](https://en.wikipedia.org/wiki/Spigot_algorithm)
// time complexity: O(n)
// space complexity: O(n)
// author(s) [red_byte](https://github.com/i-redbyte)
// see spigotpi_test.go

package pi

import "strconv"

func Spigot(n int) string {
	pi := ""
	boxes := n * 10 / 3
	remainders := make([]int, boxes)
	for i := 0; i < boxes; i++ {
		remainders[i] = 2
	}
	digitsHeld := 0
	for i := 0; i < n; i++ {
		carriedOver := 0
		sum := 0
		for j := boxes - 1; j >= 0; j-- {
			remainders[j] *= 10
			sum = remainders[j] + carriedOver
			quotient := sum / (j*2 + 1)
			remainders[j] = sum % (j*2 + 1)
			carriedOver = quotient * j
		}
		remainders[0] = sum % 10
		q := sum / 10
		switch q {
		case 9:
			digitsHeld++
		case 10:
			q = 0
			for k := 1; k <= digitsHeld; k++ {
				replaced, _ := strconv.Atoi(pi[i-k : i-k+1])
				if replaced == 9 {
					replaced = 0
				} else {
					replaced++
				}
				pi = delChar(pi, i-k)
				pi = pi[:i-k] + strconv.Itoa(replaced) + pi[i-k:]
			}
			digitsHeld = 1
		default:
			digitsHeld = 1
		}
		pi += strconv.Itoa(q)
	}
	return pi
}

func delChar(s string, index int) string {
	tmp := []rune(s)
	return string(append(tmp[0:index], tmp[index+1:]...))
}
