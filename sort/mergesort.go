package sort

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/min"
)

func merge[T constraints.Ordered](a []T, b []T) []T {

	var r = make([]T, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r

}

// Merge Perform merge sort on a slice
func Merge[T constraints.Ordered](items []T) []T {

	if len(items) < 2 {
		return items

	}

	var middle = len(items) / 2
	var a = Merge(items[:middle])
	var b = Merge(items[middle:])
	return merge(a, b)

}

func MergeIter[T constraints.Ordered](items []T) []T {
	for step := 1; step < len(items); step += step {
		for i := 0; i+step < len(items); i += 2 * step {
			tmp := merge(items[i:i+step], items[i+step:min.Int(i+2*step, len(items))])
			copy(items[i:], tmp)
		}
	}
	return items
}

// MergeParallel Perform merge sort on a slice using goroutines
func ParallelMerge[T constraints.Ordered](items []T) []T {

	if len(items) < 2 {
		return items

	}

	if len(items) < 2048 {
		return Merge(items)
	}

	done := make(chan struct{})
	defer close(done)

	var middle = len(items) / 2
	var a []T
	go func() {
		a = MergeParallel(items[:middle])
		done <- struct{}{}
	}()
	var b = MergeParallel(items[middle:])
	<-done
	return merge(a, b)
}
