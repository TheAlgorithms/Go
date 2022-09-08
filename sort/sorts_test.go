package sort_test

import (
	"reflect"
	"testing"

	"github.com/TheAlgorithms/Go/sort"
)

func testInt(t *testing.T, sortingFunction func([]int) []int) {
	sortTestsInt := []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Unsigned int",
		},
		//Reversed slice
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Unsigned int",
		},
		//Sorted slice
		{
			input:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Signed int",
		},
		//Reversed slice
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Signed int",
		},
		//Reversed slice, even length
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Signed #2 int",
		},
		//Random order with repetitions
		{
			input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
			name:     "Random order Signed int",
		},
		//Single-entry slice
		{
			input:    []int{1},
			expected: []int{1},
			name:     "Singleton int",
		},
		// Empty slice
		{
			input:    []int{},
			expected: []int{},
			name:     "Empty Slice int",
		},
	}
	for _, test := range sortTestsInt {
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

func testFloats(t *testing.T, sortingFunction func([]float32) []float32) {
	sortTestsFloat := []struct {
		input    []float32
		expected []float32
		name     string
	}{
		//Sorted slice
		{
			input:    []float32{1.23, 2.14, 2.34, 4.67, 5.12, 6.51, 7.09, 7.15, 8.45, 9.13},
			expected: []float32{1.23, 2.14, 2.34, 4.67, 5.12, 6.51, 7.09, 7.15, 8.45, 9.13},
			name:     "Sorted Unsigned float",
		},
		//Reversed slice
		{
			input:    []float32{6.76, 5.61, 3.98, 2.11},
			expected: []float32{2.11, 3.98, 5.61, 6.76},
			name:     "Reversed Unsigned float",
		},
		//Sorted slice
		{
			input:    []float32{-9.14, -7.45, -3.12, 1.23, 5.67},
			expected: []float32{-9.14, -7.45, -3.12, 1.23, 5.67},
			name:     "Sorted Signed float",
		},
		//Reversed slice
		{
			input:    []float32{5.67, 1.23, -3.12, -7.45, -9.14},
			expected: []float32{-9.14, -7.45, -3.12, 1.23, 5.67},
			name:     "Reversed Signed float",
		},
		//Random order with repetitions
		{
			input:    []float32{-5.52, 3.14, -1.23, -9.34, 7.6, 3.14, 2.12, -3.45},
			expected: []float32{-9.34, -5.52, -3.45, -1.23, 2.12, 3.14, 3.14, 7.6},
			name:     "Random order Signed float",
		},
		//Single-entry slice
		{
			input:    []float32{1.23},
			expected: []float32{1.23},
			name:     "Singleton float",
		},
		// Empty slice
		{
			input:    []float32{},
			expected: []float32{},
			name:     "Empty Slice float",
		},
	}
	for _, test := range sortTestsFloat {
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

func testStrings(t *testing.T, sortingFunction func([]string) []string) {
	sortTestsString := []struct {
		input    []string
		expected []string
		name     string
	}{
		//Sorted slice
		{
			input:    []string{"a", "b", "c", "d"},
			expected: []string{"a", "b", "c", "d"},
			name:     "Sorted Chars",
		},
		//Reversed slice
		{
			input:    []string{"d", "c", "b", "a"},
			expected: []string{"a", "b", "c", "d"},
			name:     "Reversed Chars",
		},
		//Sorted slice
		{
			input:    []string{"aab", "bbc", "dab"},
			expected: []string{"aab", "bbc", "dab"},
			name:     "Sorted String",
		},
		//Reversed slice
		{
			input:    []string{"dab", "bbc", "aab"},
			expected: []string{"aab", "bbc", "dab"},
			name:     "Reversed String",
		},
		//Random order with repetitions
		{
			input:    []string{"tea", "coffee", "books", "ant", "belt", "books"},
			expected: []string{"ant", "belt", "books", "books", "coffee", "tea"},
			name:     "Random order Strings",
		},
		//Single-entry slice
		{
			input:    []string{"hello"},
			expected: []string{"hello"},
			name:     "Singleton string",
		},
		// Empty slice
		{
			input:    []string{},
			expected: []string{},
			name:     "Empty Slice string",
		},
	}
	for _, test := range sortTestsString {
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
	testInt(t, sort.Bubble[int])
	testStrings(t, sort.Bubble[string])
	testFloats(t, sort.Bubble[float32])
}

func TestExchange(t *testing.T) {
	testInt(t, sort.Exchange[int])
	testStrings(t, sort.Exchange[string])
	testFloats(t, sort.Exchange[float32])
}

func TestInsertion(t *testing.T) {
	testInt(t, sort.Insertion[int])
	testStrings(t, sort.Insertion[string])
	testFloats(t, sort.Insertion[float32])
}

func TestMerge(t *testing.T) {
	testInt(t, sort.Merge[int])
	testStrings(t, sort.Merge[string])
	testFloats(t, sort.Merge[float32])
}

func TestMergeIter(t *testing.T) {
	testInt(t, sort.MergeIter[int])
	testStrings(t, sort.MergeIter[string])
	testFloats(t, sort.MergeIter[float32])
}

func TestHeap(t *testing.T) {
	testInt(t, sort.HeapSort)
}

func TestCount(t *testing.T) {
	testInt(t, sort.Count[int])
}

func TestQuick(t *testing.T) {
	testInt(t, sort.Quicksort[int])
	testStrings(t, sort.Quicksort[string])
	testFloats(t, sort.Quicksort[float32])
}

func TestShell(t *testing.T) {
	testInt(t, sort.Shell[int])
	testStrings(t, sort.Shell[string])
	testFloats(t, sort.Shell[float32])
}

func TestRadix(t *testing.T) {
	testInt(t, sort.RadixSort)
}

func TestSimple(t *testing.T) {
	testInt(t, sort.Simple[int])
	testStrings(t, sort.Simple[string])
	testFloats(t, sort.Simple[float32])
}

func TestImprovedSimple(t *testing.T) {
	testInt(t, sort.ImprovedSimple[int])
	testStrings(t, sort.ImprovedSimple[string])
	testFloats(t, sort.ImprovedSimple[float32])
}

func TestSelection(t *testing.T) {
	testInt(t, sort.Selection[int])
	testStrings(t, sort.Selection[string])
	testFloats(t, sort.Selection[float32])
}

func TestComb(t *testing.T) {
	testInt(t, sort.Comb[int])
	testStrings(t, sort.Comb[string])
	testFloats(t, sort.Comb[float32])
}

func TestPigeonhole(t *testing.T) {
	testInt(t, sort.Pigeonhole)
}

func TestPatience(t *testing.T) {
	testInt(t, sort.Patience[int])
	testStrings(t, sort.Patience[string])
	testFloats(t, sort.Patience[float32])
}

//END TESTS

func benchmarkFramework(b *testing.B, f func(arr []int) []int) {
	var sortTestsInt = []struct {
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
		for _, test := range sortTestsInt {
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
	benchmarkFramework(b, sort.HeapSort)
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
