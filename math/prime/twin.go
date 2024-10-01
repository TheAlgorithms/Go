// twin.go
// description: Returns Twin Prime of n
// details:
// For any integer n, twin prime is (n + 2)
// if and only if both n and (n + 2) both are prime
// wikipedia: https://en.wikipedia.org/wiki/Twin_prime
// time complexity: O(log n)
// space complexity: O(1)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see twin_test.go

package prime

// This function returns twin prime for given number
// returns (n + 2) if both n and (n + 2) are prime
// -1 otherwise
func Twin(n int) (int, bool) {
	if OptimizedTrialDivision(int64(n)) && OptimizedTrialDivision(int64(n+2)) {
		return n + 2, true
	}
	return -1, false
}
