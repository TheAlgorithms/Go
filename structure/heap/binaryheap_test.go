package heap

import (
	"fmt"
	"testing"
)

type MyInt32 int32

func (a MyInt32) PriorTo(b MyInt32) bool {
	return a < b
}
func main() {
	h := NewBinaryHeapFromSlice[MyInt32, int32]([]MyInt32{4, 2})
	h.Push(3)
	fmt.Println(h.Pop())
	testing.Benchmark(nil)
}
