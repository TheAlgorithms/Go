package fenwicktree

// package main

import (
	"errors"
)

// error message for index out of bound
var (
	errorIndex = errors.New("Index out of bound")
)

// type for slice, where fenwick logic is implemented
type fenwicktree []int

// checker for validation of index
func (arr fenwicktree) indexChecker(index int) bool {
	return (index < len(arr)) && (index >= 0)
}

// add value on specific index of the slice
func (arr fenwicktree) addVal(index int, val int) error {

	if !arr.indexChecker(index) {
		return errorIndex
	}

	for i := index; i < len(arr); i = i | (i + 1) {
		arr[i] += val
	}

	return nil
}

// returns the sum of slice values from 0 to index
func (arr fenwicktree) prefixSum(index int) (int, error) {
	if !arr.indexChecker(index) {
		return 0, errorIndex
	}

	sum := 0

	for i := index; i >= 0; i = (i & (i + 1)) - 1 {
		sum += arr[i]
	}

	return sum, nil
}

// returns the sum of slice values from start index to ending index
func (arr fenwicktree) intervalSum(startInd int, endInd int) (int, error) {

	endSum, err := arr.prefixSum(endInd)
	if err != nil {
		return 0, err
	}
	startSum, err := arr.prefixSum(startInd - 1)
	if err != nil {
		return 0, err
	}

	// a[l]+...+a[r] = (a[1]+...+a[r]) - (a[1]+...+a[l-1])
	return endSum - startSum, nil
}

// returns the value at specific index
func (arr fenwicktree) getElem(ind int) (int, error) {
	return arr.intervalSum(ind, ind)
}

/*
func main() {
	fen := make(fenwicktree, 5)
	fen.addVal(1, 3)
	fen.addVal(2, 1)

	sum, err := fen.prefixSum(3)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fen.addVal(0, 4)
	val, err := fen.getElem(2)

	interSum, err := fen.intervalSum(1, 4)

	fmt.Printf("%d %d %d", sum, val, interSum)
}

*/
