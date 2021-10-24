package min

// Int is a function which returns the minimum of all the integers provided as arguments.
func Int(values ...int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
