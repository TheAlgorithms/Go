// fibonacci.go
// description: Implementation of the Fibonacci sequence using dynamic programming
// time complexity: O(n)
// space complexity: O(1)
package dynamic

// https://www.geeksforgeeks.org/program-for-nth-fibonacci-number/

// NthFibonacci returns the nth Fibonacci Number
func NthFibonacci(n uint) uint {
	if n == 0 {
		return 0
	}

	// n1 and n2 are the (i-1)th and ith Fibonacci numbers, respectively
	var n1, n2 uint = 0, 1

	for i := uint(1); i < n; i++ {
		n3 := n1 + n2
		n1 = n2
		n2 = n3
	}

	return n2
}
