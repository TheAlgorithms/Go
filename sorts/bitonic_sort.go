/*
   This algorithm works only with lengths of the power of two.
   A Bitonic Sort is a parallel comparison-based sorting algorithm with does O(n Log 2n)
   comparisons, it is based on converting a sequence into a BitonicSequence.

   A Bitonic Sequence is a sequence of numbers in which the values of the numbers
   in the first half are strictly increasing and the values in the second half are strictly decreasing,
   so the values start growing and then go down. For example, 6 7 8 9 5 3 2 1.
   Any of those two parts, increasing or decreasing, can be empty.

   To turn a sequence into a BitonicSequence, it is recursively split into two parts,
   the first one is sorted increasingly (direction true), while the second part
   is sorted decreasingly (direction false).

   The algorithm starts sorting a group of four, the first two elements increasingly,
   and the second two elements decreasingly. The next step if sorting a group of eight,
   the first four elements increasingly and the second two elements decreasingly. And so on,
   until a group of the same size as the original array is considered the first group,
   and consequently, sorted increasingly.
*/

package main

func bitonicSort(sequence []int) []int {
	//true for increasing order, false for decreasing
	return bitonicSortImplementation(true, sequence)
}

func bitonicSortImplementation(direction bool, sequence []int) []int {
	sequenceLength := len(sequence)

	if sequenceLength <= 1 {
		return sequence
	} else {
		firstHalf := bitonicSortImplementation(true, sequence[0:(sequenceLength/2)])
		secondHalf := bitonicSortImplementation(false, sequence[(sequenceLength/2):sequenceLength])

		return bitonicMerge(direction, append(firstHalf, secondHalf...))
	}
}

func bitonicMerge(direction bool, sequence []int) []int {
	sequenceLength := len(sequence)

	if sequenceLength <= 1 {
		return sequence
	} else {
		bitonicCompare(direction, sequence)

		firstHalf := bitonicMerge(direction, sequence[0:(sequenceLength/2)])
		secondHalf := bitonicMerge(direction, sequence[(sequenceLength/2):sequenceLength])

		return append(firstHalf, secondHalf...)

	}
}

func bitonicCompare(direction bool, sequence []int) []int {
	comparisonDistance := len(sequence) / 2

	for i := 0; i < comparisonDistance; i++ {
		if sequence[i] > sequence[i+comparisonDistance] == direction {
			aux := sequence[i]
			sequence[i] = sequence[i+comparisonDistance]
			sequence[i+comparisonDistance] = aux
		}
	}

	return sequence
}
