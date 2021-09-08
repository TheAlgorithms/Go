package max

// Int gives max of two integers
// defining it here to be used in other subpackages of the repo.
func Int(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
