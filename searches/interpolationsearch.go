package searches

func InterpolationSearch(arr []int, x int) int {
	var first = 0
	var last = len(arr) - 1

	for first <= last && x >= arr[first] && x <= arr[last] {
		if first == last {
			if arr[first] == x {
				return first
			}
			return -1
		}

		tmp := first + int(((last-first)/(arr[last]-arr[first]))*(x-arr[first]))

		if arr[tmp] == x {
			return tmp
		}

		if arr[tmp] < x {
			first = tmp + 1
		} else {
			last = tmp - 1
		}
	}

	return -1
}
