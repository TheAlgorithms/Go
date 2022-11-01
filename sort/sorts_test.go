package sort_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/TheAlgorithms/Go/sort"
)

func testFramework(t *testing.T, sortingFunction func([]int) []int) {
	sortTests := []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Unsigned",
		},
		//Reversed slice
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Unsigned",
		},
		//Sorted slice
		{
			input:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Signed",
		},
		//Reversed slice
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Signed",
		},
		//Reversed slice, even length
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Signed #2",
		},
		//Random order with repetitions
		{
			input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
			name:     "Random order Signed",
		},
		//Single-entry slice
		{
			input:    []int{1},
			expected: []int{1},
			name:     "Singleton",
		},
		// Empty slice
		{
			input:    []int{},
			expected: []int{},
			name:     "Empty Slice",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := sortingFunction(test.input)
			sorted := reflect.DeepEqual(actual, test.expected)
			if !sorted {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}
}

func testFrameworkWithStrings(t *testing.T, sortingFunction func([]string) []string) {
	sortTests := []struct {
		input    []string
		expected []string
		name     string
	}{
		//Sorted slice
		{
			input:    []string{"1", "10", "2", "3", "4", "5", "6", "7", "8", "9"},
			expected: []string{"1", "10", "2", "3", "4", "5", "6", "7", "8", "9"},
			name:     "Sorted String slice",
		},
		//Reversed slice
		{
			input:    []string{"9", "7", "6", "5", "4", "3", "2", "10", "1"},
			expected: []string{"1", "10", "2", "3", "4", "5", "6", "7", "9"},
			name:     "Reversed String slice",
		},
		//Alphanumeric slice
		{
			input:    []string{"a1", "b2", "b3", "b1", "c1", "c5", "cc", "c4", "a4", "d22", "e12", "12r", "9f", "f", "t", "a", "g5"},
			expected: []string{"12r", "9f", "a", "a1", "a4", "b1", "b2", "b3", "c1", "c4", "c5", "cc", "d22", "e12", "f", "g5", "t"},
			name:     "Alphanumeric String slice",
		},
		//Strings with repetitions
		{
			input:    []string{"a", "b2", "b2", "b1", "c1", "c5", "cc", "c4", "a", "d22", "e12", "12r", "9f", "e12", "t", "a", "g5"},
			expected: []string{"12r", "9f", "a", "a", "a", "b1", "b2", "b2", "c1", "c4", "c5", "cc", "d22", "e12", "e12", "g5", "t"},
			name:     "Repetitive String",
		},
		//Single-entry slice
		{
			input:    []string{"a"},
			expected: []string{"a"},
			name:     "Singleton",
		},
		// Empty slice
		{
			input:    []string{},
			expected: []string{},
			name:     "Empty Slice",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := sortingFunction(test.input)
			sorted := reflect.DeepEqual(actual, test.expected)
			if !sorted {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}
}

//BEGIN TESTS

func TestBubble(t *testing.T) {
	testFramework(t, sort.Bubble[int])
	testFrameworkWithStrings(t, sort.Bubble[string])
}

func TestExchange(t *testing.T) {
	testFramework(t, sort.Exchange[int])
	testFrameworkWithStrings(t, sort.Exchange[string])
}

func TestInsertion(t *testing.T) {
	testFramework(t, sort.Insertion[int])
	testFrameworkWithStrings(t, sort.Insertion[string])
}

func TestMerge(t *testing.T) {
	testFramework(t, sort.Merge[int])
	testFrameworkWithStrings(t, sort.Merge[string])
}

func TestMergeIter(t *testing.T) {
	testFramework(t, sort.MergeIter[int])
	testFrameworkWithStrings(t, sort.MergeIter[string])
}

func TestMergeParallel(t *testing.T) {
	testFramework(t, sort.ParallelMerge[int])
	testFrameworkWithStrings(t, sort.ParallelMerge[string])

	// Test parallel merge sort with a large slice
	t.Run("ParallelMerge on large slice", func(t *testing.T) {
		rand.Seed(time.Now().UnixNano())
		size := 100000
		randomLargeSlice := make([]int, size)
		for i := range randomLargeSlice {
			randomLargeSlice[i] = rand.Intn(size)
		}
		sortedSlice := sort.ParallelMerge[int](randomLargeSlice)
		for i := 0; i < len(sortedSlice)-1; i++ {
			if sortedSlice[i] > sortedSlice[i+1] {
				t.Errorf("ParallelMerge failed")
			}
		}
	})
}

func TestHeap(t *testing.T) {
	testFramework(t, sort.HeapSort[int])
	testFrameworkWithStrings(t, sort.HeapSort[string])
}

func TestCount(t *testing.T) {
	testFramework(t, sort.Count[int])
}

func TestQuick(t *testing.T) {
	testFramework(t, sort.Quicksort[int])
	testFrameworkWithStrings(t, sort.Quicksort[string])
}

func TestShell(t *testing.T) {
	testFramework(t, sort.Shell[int])
	testFrameworkWithStrings(t, sort.Shell[string])
}

func TestRadix(t *testing.T) {
	testFramework(t, sort.RadixSort[int])
}

func TestSimple(t *testing.T) {
	testFramework(t, sort.Simple[int])
	testFrameworkWithStrings(t, sort.Simple[string])
}

func TestImprovedSimple(t *testing.T) {
	testFramework(t, sort.ImprovedSimple[int])
	testFrameworkWithStrings(t, sort.ImprovedSimple[string])
}

func TestSelection(t *testing.T) {
	testFramework(t, sort.Selection[int])
	testFrameworkWithStrings(t, sort.Selection[string])
}

func TestComb(t *testing.T) {
	testFramework(t, sort.Comb[int])
	testFrameworkWithStrings(t, sort.Comb[string])
}

func TestPigeonhole(t *testing.T) {
	testFramework(t, sort.Pigeonhole[int])
}

func TestPatience(t *testing.T) {
	testFramework(t, sort.Patience[int])
	testFrameworkWithStrings(t, sort.Patience[string])
}

//END TESTS

func benchmarkFramework(b *testing.B, f func(arr []int) []int) {
	var sortTests = []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted Unsigned"},
		//Reversed slice
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Unsigned"},
		//Sorted slice
		{[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted Signed"},
		//Reversed slice
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Signed"},
		//Reversed slice, even length
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Signed #2"},
		//Random order with repetitions
		{[]int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10}, "Random order Signed"},
		//Empty slice
		{[]int{}, []int{}, "Empty"},
		//Single-entry slice
		{[]int{1}, []int{1}, "Singleton"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			f(test.input)
		}
	}
}

//BEGIN BENCHMARKS

func BenchmarkBubble(b *testing.B) {
	benchmarkFramework(b, sort.Bubble[int])
}

func BenchmarkExchange(b *testing.B) {
	benchmarkFramework(b, sort.Exchange[int])
}

func BenchmarkInsertion(b *testing.B) {
	benchmarkFramework(b, sort.Insertion[int])
}

func BenchmarkMerge(b *testing.B) {
	benchmarkFramework(b, sort.Merge[int])
}

func BenchmarkMergeIter(b *testing.B) {
	benchmarkFramework(b, sort.MergeIter[int])
}

func BenchmarkMergeParallel(b *testing.B) {
	benchmarkFramework(b, sort.ParallelMerge[int])
}

func BenchmarkHeap(b *testing.B) {
	benchmarkFramework(b, sort.HeapSort[int])
}

func BenchmarkCount(b *testing.B) {
	benchmarkFramework(b, sort.Count[int])
}

func BenchmarkQuick(b *testing.B) {
	benchmarkFramework(b, sort.Quicksort[int])
}

func BenchmarkShell(b *testing.B) {
	benchmarkFramework(b, sort.Shell[int])
}

func BenchmarkRadix(b *testing.B) {
	benchmarkFramework(b, sort.RadixSort[int])
}

func BenchmarkSimple(b *testing.B) {
	benchmarkFramework(b, sort.Simple[int])
}

func BenchmarkImprovedSimple(b *testing.B) {
	benchmarkFramework(b, sort.ImprovedSimple[int])
}

// Very Slow, consider commenting
func BenchmarkSelection(b *testing.B) {
	benchmarkFramework(b, sort.Selection[int])
}

func BenchmarkComb(b *testing.B) {
	benchmarkFramework(b, sort.Comb[int])
}

func BenchmarkPigeonhole(b *testing.B) {
	benchmarkFramework(b, sort.Pigeonhole[int])
}

func BenchmarkPatience(b *testing.B) {
	benchmarkFramework(b, sort.Patience[int])
}
