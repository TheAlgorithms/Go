package circularqueue

import "testing"

func TestCircularQueue(t *testing.T) {
	t.Run("Size Check", func(t *testing.T) {
		_, err := NewCircularQueue[int](-3)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		queue, _ := NewCircularQueue[int](5)
		expectedSize := 5
		gotSize := queue.Size()
		if gotSize != expectedSize {
			t.Errorf("Expected size: %v, got: %v\n", expectedSize, gotSize)
		}

		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Enqueue(4)
		queue.Enqueue(5)

		err = queue.Enqueue(6)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		expectedSize = 5
		gotSize = queue.Size()
		if gotSize != expectedSize {
			t.Errorf("Expected size: %v, got: %v\n", expectedSize, gotSize)
		}

		queue.Dequeue()
		queue.Dequeue()

		err = queue.Enqueue(6)
		if err != nil {
			t.Errorf("Expected nil, got error: %v\n", err.Error())
		}

		expectedSize = 5
		gotSize = queue.Size()
		if gotSize != expectedSize {
			t.Errorf("Expected size: %v, got: %v\n", expectedSize, gotSize)
		}
	})
	t.Run("Enqueue", func(t *testing.T) {
		queue, _ := NewCircularQueue[int](10)

		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)

		expected := 1
		got, err := queue.Peek()
		if err != nil {
			t.Error(err.Error())
		}
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}
	})
	t.Run("Dequeue", func(t *testing.T) {
		queue, _ := NewCircularQueue[string](10)

		queue.Enqueue("one")
		queue.Enqueue("two")
		queue.Enqueue("three")

		expected := "one"
		got, err := queue.Dequeue()
		if err != nil {
			t.Error(err.Error())
		}
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}

		expected = "two"
		got, err = queue.Peek()
		if err != nil {
			t.Error(err.Error())
		}
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}
	})
	t.Run("Circularity", func(t *testing.T) {
		queue, _ := NewCircularQueue[int](10)

		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Dequeue()
		queue.Dequeue()
		queue.Enqueue(4)
		queue.Enqueue(5)
		queue.Dequeue()

		expected := 4
		got, err := queue.Peek()
		if err != nil {
			t.Error(err.Error())
		}
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}
	})
	t.Run("IsFull", func(t *testing.T) {
		queue, _ := NewCircularQueue[bool](2)
		queue.Enqueue(false)
		queue.Enqueue(true)

		expected := true
		got := queue.IsFull()
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}

		queue.Dequeue()
		queue.Dequeue()

		expected = false
		got = queue.IsFull()
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}
	})
	t.Run("IsEmpty", func(t *testing.T) {
		queue, _ := NewCircularQueue[float64](2)

		expected := true
		got := queue.IsEmpty()
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}

		queue.Enqueue(1.0)

		expected = false
		got = queue.IsEmpty()
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}

	})
	t.Run("Peak", func(t *testing.T) {
		queue, _ := NewCircularQueue[rune](10)

		queue.Enqueue('a')
		queue.Enqueue('b')
		queue.Enqueue('c')

		expected := 'a'
		got, err := queue.Peek()
		if err != nil {
			t.Error(err.Error())
		}

		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}
	})
}

// BenchmarkCircularQueue benchmarks the CircularQueue implementation.
func BenchmarkCircularQueue(b *testing.B) {
	b.Run("Enqueue", func(b *testing.B) {
		queue, _ := NewCircularQueue[int](1000)
		for i := 0; i < b.N; i++ {
			queue.Enqueue(i)
		}
	})

	b.Run("Dequeue", func(b *testing.B) {
		queue, _ := NewCircularQueue[int](1000)
		for i := 0; i < 1000; i++ {
			queue.Enqueue(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			queue.Dequeue()
		}
	})

	b.Run("Peek", func(b *testing.B) {
		queue, _ := NewCircularQueue[int](1000)
		for i := 0; i < 1000; i++ {
			queue.Enqueue(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			queue.Peek()
		}
	})
}
