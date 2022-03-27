// millerrabinprimalitytest.go
// description: An implementation of Miller-Rabin primality test
// details:
// A simple implementation of Miller-Rabin Primality Test
// [Miller-Rabin primality test Wiki](https://en.wikipedia.org/wiki/Miller–Rabin_primality_test)
// author(s) [Taj](https://github.com/tjgurwara99)
// see millerrabinprimalitytest_test.go

package prime

import (
	"math/rand"

	"github.com/TheAlgorithms/Go/math/modular"
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

func MillerTest(num, witness int64) (bool, error) {
	d, _ := findRD(num)
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

// MillerTest This is the intermediate step that repeats within the
// miller rabin primality test for better probabilitic chances of
// receiving the correct result.
func MillerRandomTest(num int64) (bool, error) {
	random := rand.Int63n(num-1) + 2
	return MillerTest(num, random)
}

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
	

// MillerRabinProbabilistic is a probabilistic test for primality of an integer based of the algorithm devised by Miller and Rabin.
func MillerRabinProbabilistic(num, rounds int64) (bool, error) {
	if num <= 4 {
		if num == 2 || num == 3 {
			return true, nil
		}
		return false, nil
	}
	if num%2 == 0 {
		return false, nil
	}

	for i := int64(0); i < rounds; i++ {
		val, err := MillerRandomTest(d, num)
		if err != nil {
			return false, err
		}
		if !val {
			return false, nil
		}
	}
	return true, nil
}

func MillerRabinDeterministic(num int64) (bool, error) {
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
	
	switch {
	case num < 1_373_653:
		return MillerTestMultiple(num, 2, 3)
	case num < 25_326_001:
		return MillerTestMultiple(num, 2, 3, 5)
	case num < 1_122_004_669_633:
		return MillerTestMultiple(num, 2, 13, 23, 1_662_803)
	default:
		return MillerTestMultiple(num, 2, 3, 5, 7, 11, 13, 17, 19, 23, 31, 37)
	}
}
