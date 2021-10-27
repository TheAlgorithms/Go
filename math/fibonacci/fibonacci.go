// fibonacci.go
// description: Calculates the nth fibonacci number
// author(s) [jyothsnashaji](https://github.com/jyothsnashaji)

package fibonacci

func Nthfib(n int) int {
	if n < 1 {
		return 0
	} else if n <= 2 {
		return 1
	}
	var first int = 1
	var second int = 1
	var result int = first + second
	for i := 3; i <= n; i++ {
		result = first + second
		first = second
		second = result
	}
	return result
}
