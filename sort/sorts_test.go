package sort_test

import (
	"reflect"
	"testing"

	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/sort"
)

type testData[T constraints.Ordered] struct {
	input    []T
	expected []T
	name     string
}

var intTestData = []testData[int]{
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

var floatTestData = []testData[float32]{
	//Sorted slice
	{
		input:    []float32{1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0},
		expected: []float32{1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0},
		name:     "Sorted float32",
	},
	//Reversed slice
	{
		input:    []float32{2.0, 1.9, 1.8, 1.7, 1.6, 1.5, 1.4, 1.3, 1.2, 1.1},
		expected: []float32{1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0},
		name:     "Reversed float32",
	},
	//Sorted slice
	{
		input:    []float32{-2.0, -1.9, -1.8, -1.7, -1.6, -1.5, -1.4, -1.3, -1.2, -1.1, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0},
		expected: []float32{-2.0, -1.9, -1.8, -1.7, -1.6, -1.5, -1.4, -1.3, -1.2, -1.1, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0},
		name:     "Sorted float32",
	},
	//Reversed slice
	{
		input:    []float32{2.0, 1.9, 1.8, 1.7, 1.6, 1.5, 1.4, 1.3, 1.2, 1.1, 0, -1.1, -1.2, -1.3, -1.4, -1.5, -1.6, -1.7, -1.8, -1.9, -2.0},
		expected: []float32{-2.0, -1.9, -1.8, -1.7, -1.6, -1.5, -1.4, -1.3, -1.2, -1.1, 0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0},
		name:     "Reversed float32",
	},
	//Reversed slice, even length
	{
		input:    []float32{2.0, 1.9, 1.8, 1.7, 1.6, 1.5, 1.4, 1.3, 1.2, 1.1, -1.1, -1.2, -1.3, -1.4, -1.5, -1.6, -1.7, -1.8, -1.9, -2.0},
		expected: []float32{-2.0, -1.9, -1.8, -1.7, -1.6, -1.5, -1.4, -1.3, -1.2, -1.1, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0},
		name:     "Reversed float32 #2",
	},
	//Random order with repetitions
	{
		input:    []float32{-1.5, 1.7, 1.4, -1.2, 1.6, 1.5, 1.8, 1.3, 1.2, -1.7, -1.1, 0, -1.3, 1.9, -1.6, -1.4, 2, 1.9, 1.1, -1.8, -1.9, -2},
		expected: []float32{-2.0, -1.9, -1.8, -1.7, -1.6, -1.5, -1.4, -1.3, -1.2, -1.1, 0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 1.9, 2},
		name:     "Random order float32",
	},
	//Single-entry slice
	{
		input:    []float32{1.0},
		expected: []float32{1.0},
		name:     "Singleton",
	},
	// Empty slice
	{
		input:    []float32{},
		expected: []float32{},
		name:     "Empty Slice",
	},
}

func testFramework(t *testing.T, sortingFunction func([]int) []int) {
	for _, test := range intTestData {
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

func genericTestFramework[T constraints.Ordered](t *testing.T, data []testData[T], sortingFunction func([]T) []T) {
	for _, test := range data {
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
}

func TestExchange(t *testing.T) {
	testFramework(t, sort.Exchange[int])
}

func TestInsertion(t *testing.T) {
	testFramework(t, sort.Insertion[int])
}

func TestMerge(t *testing.T) {
	testFramework(t, sort.Merge[int])
}

func TestMergeIter(t *testing.T) {
	testFramework(t, sort.MergeIter[int])
}

func TestHeap(t *testing.T) {
	genericTestFramework(t, intTestData, sort.HeapSort[int])
	genericTestFramework(t, floatTestData, sort.HeapSort[float32])
}

func TestCount(t *testing.T) {
	testFramework(t, sort.Count[int])
}

func TestQuick(t *testing.T) {
	testFramework(t, sort.Quicksort[int])
}

func TestShell(t *testing.T) {
	testFramework(t, sort.Shell[int])
}

func TestRadix(t *testing.T) {
	testFramework(t, sort.RadixSort)
}

func TestSimple(t *testing.T) {
	testFramework(t, sort.Simple[int])
}

func TestImprovedSimple(t *testing.T) {
	testFramework(t, sort.ImprovedSimple[int])
}

func TestSelection(t *testing.T) {
	testFramework(t, sort.Selection[int])
}

func TestComb(t *testing.T) {
	testFramework(t, sort.Comb[int])
}

func TestPigeonhole(t *testing.T) {
	testFramework(t, sort.Pigeonhole)
}

func TestPatience(t *testing.T) {
	testFramework(t, sort.Patience[int])
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
	benchmarkFramework(b, sort.RadixSort)
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
	benchmarkFramework(b, sort.Pigeonhole)
}

func BenchmarkPatience(b *testing.B) {
	benchmarkFramework(b, sort.Patience[int])
}
