// mode_test.go
// author(s): [CalvinNJK] (https://github.com/CalvinNJK)
// description: Finding Mode Value In an Array
// see mode.go

package mode

func Mode(numbers []float64) float64 {

	// Find the size of the array
	n := len(numbers)

	// declare variables
	maxCount := 0
	maxValue := 0.0

	// Loop for the whole array
	for i := 0; i < n; i++ {

		count := 0

		// check the selected Value have any repeated same value
		for k := 0; k < n; k++ {

			// If found repeated value, add into count
			if numbers[k] == numbers[i] {
				count++
			}
		}

		// Check is the selected value's count in the array,
		// is it more than previous value(MaxValue)'s count
		if count > maxCount {
			maxCount = count
			maxValue = numbers[i]
		}
	}

	return maxValue
}
