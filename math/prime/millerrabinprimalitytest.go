package prime

import (
	"math/rand"

	"github.com/TheAlgorithms/Go/math/modulararithmetic"
)

// findD accepts a number and returns the
// odd number d such that num = 2^r * d - 1
func findRD(num int64) (int64, int64) {
	r := int64(0)
	d := num - 1
	for num%2 == 0 {
		d /= 2
		r++
	}
	return d, r
}

// MillerTest This is the intermediate step that repeats within the
// miller rabin primality test for better probabilitic chances of
// receiving the correct result.
func MillerTest(d, num int64) (bool, error) {
	random := rand.Int63n(num-1) + 2

	res, err := modulararithmetic.ModularExponentiation(random, d, num)

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

// MillerRabinTest Probabilistic test for primality of an integer based of the algorithm devised by Miller and Rabin.
func MillerRabinTest(num, rounds int64) (bool, error) {
	if num <= 4 {
		if num == 2 || num == 3 {
			return true, nil
		}
		return false, nil
	}
	if num%2 == 0 {
		return false, nil
	}
	d, _ := findRD(num)

	for i := int64(0); i < rounds; i++ {
		val, err := MillerTest(d, num)
		if err != nil {
			return false, err
		}
		if !val {
			return false, nil
		}
	}
	return true, nil
}
