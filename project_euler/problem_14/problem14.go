/**
* Problem 14 - Longest Collatz sequence
* @see {@link https://projecteuler.net/problem=14}
*
* The following iterative sequence is defined for the set of positive integers:
* n → n/2 (n is even)
* n → 3n + 1 (n is odd)
*
* Using the rule above and starting with 13, we generate the following sequence:
* 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
*
* Which starting number, under one million, produces the longest chain?
*
* NOTE: Once the chain starts the terms are allowed to go above one million.
*
* @author ddaniel27
 */
package problem14

type dict map[uint64]uint64

var dictionary = dict{
	1: 1,
}

func Problem14(limit uint64) uint64 {
	for i := uint64(2); i <= limit; i++ {
		weightNextNode(i)
	}

	var max, maxWeight uint64
	for k, v := range dictionary {
		if v > maxWeight {
			max = k
			maxWeight = v
		}
	}

	return max
}

func weightNextNode(current uint64) uint64 {
	var next, weight uint64
	if current%2 == 0 {
		next = current / 2
	} else {
		next = (3 * current) + 1
	}
	if v, ok := dictionary[next]; !ok {
		weight = weightNextNode(next) + 1
	} else {
		weight = v + 1
	}
	dictionary[current] = weight
	return weight
}
