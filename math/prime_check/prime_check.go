//A primality test is an algorithm for determining whether an input number is prime.Among other fields of mathematics, it is used for cryptography.
//Unlike integer factorization, primality tests do not generally give prime factors, only stating whether the input number is prime or not.
//Source - Wikipedia https://en.wikipedia.org/wiki/Primality_test
package main

func NaiveApproach(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {

		if n%i == 0 {
			return false
		}
	}
	return true
}

func PairApproach(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumFunction(a int, b int) string {
	result := a + b
	return ("result of sum : " + string(result))
}

//goos: windows
//goarch: amd64
//BenchmarkNaiveApproach-12       660980034                1.81 ns/op
//BenchmarkPairApproach-12        664641356                1.82 ns/op
