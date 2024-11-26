// This file implements two versions of the Miller-Rabin primality test.
// One of the implementations is deterministic and the other is probabilistic.
// The Miller-Rabin test is one of the simplest and fastest known primality
// tests and is widely used.
// time complexity: O(k * log(n)^3)
// space complexity: O(1)
// Authors:
// [Taj](https://github.com/tjgurwara99)
// [Rak](https://github.com/raklaptudirm)

package prime

import (
	"math/rand"

	"github.com/TheAlgorithms/Go/math/modular"
)

// formatNum accepts a number and returns the
// odd number d such that num = 2^s * d + 1
func formatNum(num int64) (d int64, s int64) {
	d = num - 1
	for num%2 == 0 {
		d /= 2
		s++
	}
	return
}

// isTrivial checks if num's primality is easy to determine.
// If it is, it returns true and num's primality. Otherwise
// it returns false and false.
func isTrivial(num int64) (prime bool, trivial bool) {
	if num <= 4 {
		// 2 and 3 are primes
		prime = num == 2 || num == 3
		trivial = true
	} else {
		prime = false
		// number is trivial prime if
		// it is divisible by 2
		trivial = num%2 == 0
	}

	return
}

// MillerTest tests whether num is a strong probable prime to a witness.
// Formally: a^d ≡ 1 (mod n) or a^(2^r * d) ≡ -1 (mod n), 0 <= r <= s
func MillerTest(num, witness int64) (bool, error) {
	d, _ := formatNum(num)
	res, err := modular.Exponentiation(witness, d, num)

	if err != nil {
		return false, err
	}
	// miller conditions checks
	if res == 1 || res == num-1 {
		return true, nil
	}

	for d != num-1 {
		res = (res * res) % num
		d *= 2
		if res == 1 {
			return false, nil
		}
		if res == num-1 {
			return true, nil
		}
	}
	return false, nil
}

// MillerRandomTest This is the intermediate step that repeats within the
// miller rabin primality test for better probabilitic chances of
// receiving the correct result with random witnesses.
func MillerRandomTest(num int64) (bool, error) {
	random := rand.Int63n(num-2) + 2
	return MillerTest(num, random)
}

// MillerTestMultiple is like MillerTest but runs the test for multiple
// witnesses.
func MillerTestMultiple(num int64, witnesses ...int64) (bool, error) {
	for _, witness := range witnesses {
		prime, err := MillerTest(num, witness)
		if err != nil {
			return false, err
		}

		if !prime {
			return false, nil
		}
	}

	return true, nil
}

// MillerRabinProbabilistic is a probabilistic test for primality
// of an integer based of the algorithm devised by Miller and Rabin.
func MillerRabinProbabilistic(num, rounds int64) (bool, error) {
	if prime, trivial := isTrivial(num); trivial {
		// num is a trivial number
		return prime, nil
	}

	for i := int64(0); i < rounds; i++ {
		val, err := MillerRandomTest(num)
		if err != nil {
			return false, err
		}
		if !val {
			return false, nil
		}
	}
	return true, nil
}

// MillerRabinDeterministic is a Deterministic version of the Miller-Rabin
// test, which returns correct results for all valid int64 numbers.
func MillerRabinDeterministic(num int64) (bool, error) {
	if prime, trivial := isTrivial(num); trivial {
		// num is a trivial number
		return prime, nil
	}

	switch {
	case num < 2047:
		// witness 2 can determine the primality of any number less than 2047
		return MillerTest(num, 2)
	case num < 1_373_653:
		// witnesses 2 and 3 can determine the primality
		// of any number less than 1,373,653
		return MillerTestMultiple(num, 2, 3)
	case num < 9_080_191:
		// witnesses 31 and 73 can determine the primality
		// of any number less than 9,080,191
		return MillerTestMultiple(num, 31, 73)
	case num < 25_326_001:
		// witnesses 2, 3, and 5 can determine the
		// primality of any number less than 25,326,001
		return MillerTestMultiple(num, 2, 3, 5)
	case num < 1_122_004_669_633:
		// witnesses 2, 13, 23, and 1,662,803 can determine the
		// primality of any number less than 1,122,004,669,633
		return MillerTestMultiple(num, 2, 13, 23, 1_662_803)
	case num < 2_152_302_898_747:
		// witnesses 2, 3, 5, 7, and 11 can determine the primality
		// of any number less than 2,152,302,898,747
		return MillerTestMultiple(num, 2, 3, 5, 7, 11)
	case num < 341_550_071_728_321:
		// witnesses 2, 3, 5, 7, 11, 13, and 17 can determine the
		// primality of any number less than 341,550,071,728,321
		return MillerTestMultiple(num, 2, 3, 5, 7, 11, 13, 17)
	default:
		// witnesses 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, and 37 can determine
		// the primality of any number less than 318,665,857,834,031,151,167,461
		// which is well above the max int64 9,223,372,036,854,775,807
		return MillerTestMultiple(num, 2, 3, 5, 7, 11, 13, 17, 19, 23, 31, 37)
	}
}
