package sort

import "github.com/TheAlgorithms/Go/math/min"

func Mergesort2(items []int) []int {
	for step := 1; step < len(items); step += step {
		for i := 0; i+step < len(items); i += 2 * step {
			merge2(items, i, i+step, min.Int(i+2*step, len(items)))
		}
	}
	return items
}

func merge2(items []int, l, mid, r int) {
	l1, r1, l2, r2 := l, mid, mid, r
	var tmp []int
	for l1 < r1 && l2 < r2 {
		if items[l1] < items[l2] {
			tmp = append(tmp, items[l1])
			l1++
		} else {
			tmp = append(tmp, items[l2])
			l2++
		}
	}
	for l1 < r1 {
		tmp = append(tmp, items[l1])
		l1++
	}
	for l2 < r2 {
		tmp = append(tmp, items[l2])
		l2++
	}
	for i, v := range tmp {
		items[l+i] = v
	}
}
