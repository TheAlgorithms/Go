package sorts

import "testing"

func TestBubbleSort(t *testing.T) {
	for _, test := range sortTests {
		actual := bubbleSort(test.input)
		pos, sorted := compareSlices(actual, test.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", test.name)
			}
			t.Errorf("test %s failed at index %d", test.name, pos)
		}
	}
}

func TestCountingSort(t *testing.T) {
	for _, testcase := range sortTests {
		actual := countingSort(testcase.input)
		pos, sorted := compareSlices(actual, testcase.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", testcase.name)
			}
			t.Errorf("test %s failed at index %d", testcase.name, pos)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for _, test := range sortTests {
		actual := selectionSort(test.input)
		pos, sorted := compareSlices(actual, test.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", test.name)
			}
			t.Errorf("test %s failed at index %d", test.name, pos)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, test := range sortTests {
		actual := insertionSort(test.input)
		pos, sorted := compareSlices(actual, test.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", test.name)
			}
			t.Errorf("test %s failed at index %d", test.name, pos)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, test := range sortTests {
		actual := Mergesort(test.input)
		pos, sorted := compareSlices(actual, test.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", test.name)
			}
			t.Errorf("test %s failed at index %d", test.name, pos)
		}
	}
}

func TestHeapSort(t *testing.T) {
	for _, test := range sortTests {
		actual := heapSort(test.input)
		pos, sorted := compareSlices(actual, test.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", test.name)
			}
			t.Errorf("test %s failed at index %d", test.name, pos)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, test := range sortTests {
		actual := quickSort(test.input)
		pos, sorted := compareSlices(actual, test.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", test.name)
			}
			t.Errorf("test %s failed at index %d", test.name, pos)
		}
	}
}

func TestShellSort(t *testing.T) {
	for _, test := range sortTests {
		actual := shellSort(test.input)
		pos, sorted := compareSlices(actual, test.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", test.name)
			}
			t.Errorf("test %s failed at index %d", test.name, pos)
		}
	}
}

func TestShakerSort(t *testing.T) {
	for _, test := range sortTests {
		actual := shakersort(test.input)
		pos, sorted := compareSlices(actual, test.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", test.name)
			}
			t.Errorf("test %s failed at index %d", test.name, pos)
		}
	}
}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			bubbleSort(test.input)
		}
	}
}

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			selectionSort(test.input)
		}
	}
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			insertionSort(test.input)
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
			heapSort(test.input)
		}
	}
}

func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			quickSort(test.input)
		}
	}
}

func BenchmarkShell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			shellSort(test.input)
		}
	}
}

func BenchmarkShaker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			shakersort(test.input)
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
