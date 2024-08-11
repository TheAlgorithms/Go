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

		if err := queue.Enqueue(1); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(2); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(3); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(4); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(5); err != nil {
			t.Error(err)
		}

		err = queue.Enqueue(6)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		expectedSize = 5
		gotSize = queue.Size()
		if gotSize != expectedSize {
			t.Errorf("Expected size: %v, got: %v\n", expectedSize, gotSize)
		}

		if _, err := queue.Dequeue(); err != nil {
			t.Error(err)
		}
		if _, err := queue.Dequeue(); err != nil {
			t.Error(err)
		}

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

		if err := queue.Enqueue(1); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(2); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(3); err != nil {
			t.Error(err)
		}

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

		if err := queue.Enqueue("one"); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue("two"); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue("three"); err != nil {
			t.Error(err)
		}

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

		if err := queue.Enqueue(1); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(2); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(3); err != nil {
			t.Error(err)
		}
		if _, err := queue.Dequeue(); err != nil {
			t.Error(err)
		}
		if _, err := queue.Dequeue(); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(4); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(5); err != nil {
			t.Error(err)
		}
		if _, err := queue.Dequeue(); err != nil {
			t.Error(err)
		}

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
		if err := queue.Enqueue(false); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue(true); err != nil {
			t.Error(err)
		}

		expected := true
		got := queue.IsFull()
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}

		if _, err := queue.Dequeue(); err != nil {
			t.Error(err)
		}
		if _, err := queue.Dequeue(); err != nil {
			t.Error(err)
		}

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

		if err := queue.Enqueue(1.0); err != nil {
			t.Error(err)
		}

		expected = false
		got = queue.IsEmpty()
		if got != expected {
			t.Errorf("Expected: %v got: %v\n", expected, got)
		}

	})
	t.Run("Peak", func(t *testing.T) {
		queue, _ := NewCircularQueue[rune](10)

		if err := queue.Enqueue('a'); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue('b'); err != nil {
			t.Error(err)
		}
		if err := queue.Enqueue('c'); err != nil {
			t.Error(err)
		}

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
			if err := queue.Enqueue(i); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Dequeue", func(b *testing.B) {
		queue, _ := NewCircularQueue[int](1000)
		for i := 0; i < 1000; i++ {
			if err := queue.Enqueue(i); err != nil {
				b.Error(err)
			}
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := queue.Dequeue(); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Peek", func(b *testing.B) {
		queue, _ := NewCircularQueue[int](1000)
		for i := 0; i < 1000; i++ {
			if err := queue.Enqueue(i); err != nil {
				b.Error(err)
			}
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := queue.Peek(); err != nil {
				b.Error(err)
			}
		}
	})
}
