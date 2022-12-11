package sqrtdecompositionsimple

import (
	"testing"
)

func main() {
	slice := []int32{1, 2, 3, 4, 5, 6}
	s := NewSqrtDecompositionSimple(slice,
		func(e int32) int32 { return e },
		func(q1, q2 int32) int32 { return q1 + q2 },
		func(q, a, b int32) int32 { return q - a + b },
	)
	s.Update(3, 10)
	println(s.Query(1, 6)) // 26
	testing.Benchmark(nil)
}
