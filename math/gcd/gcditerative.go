// time complexity: O(log(min(a, b))) where a and b are the two numbers
// space complexity: O(1)

package gcd

// Iterative Faster iterative version of GcdRecursive without holding up too much of the stack
func Iterative(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
