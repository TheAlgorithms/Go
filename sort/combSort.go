// Implementation of comb sort algorithm, an improvement of bubble sort
// Reference: https://www.geeksforgeeks.org/comb-sort/

package sort

func getNextGap(gap int) int {
	gap = (gap * 10) / 13
	if gap < 1 {
		return 1
	}
	return gap
}

func combSort(data []int) []int {
	n := len(data)
	gap := n
	swapped := true

	for gap != 1 || swapped {
		gap = getNextGap(gap)
		swapped = false
		for i := 0; i < n-gap; i++ {
			if data[i] > data[i+gap] {
				data[i], data[i+gap] = data[i+gap], data[i]
				swapped = true
			}
		}
	}
	return data
}
