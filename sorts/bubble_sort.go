//Package sorts a package for demonstrating sorting algorithms in Go
package sorts
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor)-1; i++ {
			if arrayzor[i+1] < arrayzor[i] {
				arrayzor[i+1], arrayzor[i] =  arrayzor[i], arrayzor[i+1]
				swapped = true
			}
		}
	}
	return arrayzor
}
