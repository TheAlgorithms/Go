package power

// IterativePower is iterative O(logn) function for pow(x, y)
func IterativePower(n uint, power uint) uint {
	var res uint = 1
	for power > 0 {
		if (power & 1) != 0 {
			res = res * n
		}

		power = power >> 1
		n *= n
	}
	return res
}

// RecursivePower is recursive O(logn) function for pow(x, y)
func RecursivePower(n uint, power uint) uint {
	if power == 0 {
		return 1
	}
	var temp = RecursivePower(n, power/2)
	if power%2 == 0 {
		return temp * temp
	}
	return n * temp * temp
}

// RecursivePower1 is recursive O(n) function for pow(x, y)
func RecursivePower1(n uint, power uint) uint {
	if power == 0 {
		return 1
	} else if power%2 == 0 {
		return RecursivePower1(n, power/2) * RecursivePower1(n, power/2)
	} else {
		return n * RecursivePower1(n, power/2) * RecursivePower1(n, power/2)
	}
}
