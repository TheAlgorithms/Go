package heap

import "testing"

func TestInsert(t *testing.T) {
	h := New(7)
	h.Insert(5)

	if h.Max() != 5 {
		t.Errorf("max element should be 5, but found %d", h.Max())
	}
	if h.Size != 1 {
		t.Errorf("size of heap should be 1, but found %d", h.Size)
	}

	h.Insert(6)
	h.Insert(1)
	h.Insert(9)
	h.Insert(78)

	if h.Max() != 78 {
		t.Errorf("max element should be 78, but found %d", h.Max())
	}
	if h.Size != 5 {
		t.Errorf("size of heap should be 5, but found %d", h.Size)
	}
}

func TestMax(t *testing.T) {
	h := prepareHeapForTest()
	if h.Max() != 78 {
		t.Errorf("max element should be 78, but found %d", h.Max())
	}
}

func TestExtractMax(t *testing.T) {
	h := prepareHeapForTest()

	max := h.ExtractMax()
	if max != 78 {
		t.Errorf("max element should be 78, but found %d", h.Max())
	}
	if h.Size != 3 {
		t.Errorf("size of heap should be 3, but found %d", h.Size)
	}

	if h.Max() != 9 {
		t.Errorf("max element should be 9, but found %d", h.Max())
	}
}

func TestHeapify(t *testing.T) {
	size := 10
	arr := make([]int, 0, size)
	for i := 0; i < size; i++ {
		arr = append(arr, i)
	}
	h := Heapify(arr)
	for i := size - 1; i >= 0; i-- {
		max := h.ExtractMax()
		if max != i {
			t.Errorf("max element should %d, but found %d", i, max)
		}
	}
}

func prepareHeapForTest() *Heap {
	h := New(7)

	h.Insert(6)
	h.Insert(1)
	h.Insert(9)
	h.Insert(78)
	return h
}
