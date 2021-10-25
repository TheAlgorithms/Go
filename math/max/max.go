package max

// Int is a function which returns the maximum of all the integers provided as arguments.
func Int(values ...int) int {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}
