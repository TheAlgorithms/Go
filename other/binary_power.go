package binpow

import "fmt"

func power(a, n int) int {
	if (n == 0) {
		return 1
	}
	if (n % 2 == 0) {
		t := power(a, n / 2)
		return t * t
	} else {
		return power(a, n - 1) * a
	}
}

func main() {
	fmt.Println(power(2, 11))
}