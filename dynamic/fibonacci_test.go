package dynamic

import (
	"fmt"
	"testing"
)

func Test_NthFibonacci(t *testing.T) {
	// source: http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/Fibonacci/fibtable.html
	var fibonacciNumbers = []struct {
		nth       uint
		fibonacci uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{20, 6765},
		{30, 832040},
		{40, 102334155},
		{50, 12586269025},
		{60, 1548008755920},
		{70, 190392490709135},
		{80, 23416728348467685},
		{90, 2880067194370816120},
	}
	for i := range fibonacciNumbers {
		t.Run(fmt.Sprintf("the %dth Fibonacci number", fibonacciNumbers[i].nth), func(t *testing.T) {
			result := NthFibonacci(fibonacciNumbers[i].nth)
			if result != fibonacciNumbers[i].fibonacci {
				t.Errorf("Expected the %dth Fibonacci number to be %d, got %d", fibonacciNumbers[i].nth, fibonacciNumbers[i].fibonacci, result)
			}
		})
	}
}
