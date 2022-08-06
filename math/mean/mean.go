package mean

import (
	"fmt"
)

func mean(values []float64) float64{

	if len(values) <= 0 {
		fmt.Println("No values detected")
		return 0
	}

	var summation float64 = 0

	// Summation of all values in the slice
	for _,singleValue := range values{
		summation += singleValue
	}

	// Casting to float64 done for the len()
	var average = summation/float64(len(values))

	// Return of average value
	return average

}
