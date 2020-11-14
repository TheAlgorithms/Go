package sorts

import (
	"testing"
)

func testFramework(t *testing.T, sortingFunction func([]int) []int) {
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := sortingFunction(test.input)
			pos, sorted := compareSlices(actual, test.expected)
			if !sorted {
				if pos == -1 {
					t.Errorf("test %s failed due to slice length changing", test.name)
				}
				t.Errorf("test %s failed at index %d", test.name, pos)
			}
		})
	}
}

//BEGIN TESTS

func TestBubble(t *testing.T) {
	testFramework(t, bubbleSort)
}

func TestInsertion(t *testing.T) {
	testFramework(t, InsertionSort)
}

func TestMerge(t *testing.T) {
	testFramework(t, Mergesort)
}

func TestHeap(t *testing.T) {
	testFramework(t, HeapSort)
}

func TestQuick(t *testing.T) {
	testFramework(t, QuickSort)
}

func TestShell(t *testing.T) {
	testFramework(t, ShellSort)
}

func TestRadix(t *testing.T) {
	testFramework(t, RadixSort)
}

// Very slow, consider commenting
func TestSelection(t *testing.T) {
	testFramework(t, SelectionSort)
}

/* func TestTopological(t *testing.T) {
	testFramework(t, topologicalSort)
} */

//END TESTS

func benchmarkFramework(b *testing.B, sortingFunction func([]int) []int) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			sortingFunction(test.input)
		}
	}
}

func BenchmarkBubble(b *testing.B) {
	benchmarkFramework(b, bubbleSort)
}
func BenchmarkInsertion(b *testing.B) {
	benchmarkFramework(b, InsertionSort)
}
func BenchmarkMerge(b *testing.B) {
	benchmarkFramework(b, Mergesort)
}
func BenchmarkHeap(b *testing.B) {
	benchmarkFramework(b, HeapSort)
}
func BenchmarkQuick(b *testing.B) {
	benchmarkFramework(b, QuickSort)
}
func BenchmarkShell(b *testing.B) {
	benchmarkFramework(b, ShellSort)
}
func BenchmarkRadix(b *testing.B) {
	benchmarkFramework(b, RadixSort)
}

// Very Slow, consider commenting
func BenchmarkSelection(b *testing.B) {
	benchmarkFramework(b, SelectionSort)
}

func compareSlices(a []int, b []int) (int, bool) {
	if len(a) != len(b) {
		return -1, false
	}
	for pos := range a {
		if a[pos] != b[pos] {
			return pos, false
		}
	}
	return -1, true
}
