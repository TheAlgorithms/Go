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

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			bubbleSort(test.input)
		}
	}
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			InsertionSort(test.input)
		}
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			Mergesort(test.input)
		}
	}
}

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			HeapSort(test.input)
		}
	}
}

func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			QuickSort(test.input)
		}
	}
}

func BenchmarkShell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			ShellSort(test.input)
		}
	}
}

func BenchmarkRadix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			RadixSort(test.input)
		}
	}
}

// Very Slow, consider commenting
func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			SelectionSort(test.input)
		}
	}
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
