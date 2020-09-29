package main

//iterative O(logn) function for pow(x, y)
func iterativePower(n uint, power uint) uint {
	var res uint = 1
	for power > 0 {

		if (power & 1) != 0 {
			res = res * n
		}

		power = power >> 1
		n = n * n
	}
	return res
}

//recursive O(logn) function for pow(x, y)
func recursivePower(n uint, power uint) uint {
	if power == 0 {
		return 1
	}
	var temp = recursivePower(n, power/2)
	if power%2 == 0 {
		return temp * temp
	} else {
		return n * temp * temp
	}
}


//recursive O(n) function for pow(x, y)
func recursivePower1( n uint, power uint) uint{
    if power == 0 {
		return 1
	} else if power % 2 == 0 {
		return recursivePower1(n, power/2) * recursivePower1(n, power/2)
	} else
	{
		return n * recursivePower1(n, power/2) * recursivePower1(n, power/2)
	}
}
