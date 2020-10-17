// Implementation of basic radix sort algorithm
// Reference: https://en.wikipedia.org/wiki/Radix_sort

package radix

import (
	"reflect"
	"unsafe"
)

func RadixSort(a []string) {
	n := len(a)
	if n < 2 {
		return
	}
	// Put elements into a linked list.
	mem := make([]list, n)
	for i, s := range a {
		mem[i].str = s
		if i < n-1 {
			mem[i].next = &mem[i+1]
		}
	}
	res := msdRadixSort(&mem[0], n)
	for i := range a {
		a[i] = res.str
		res = res.next
	}
  return res
}
