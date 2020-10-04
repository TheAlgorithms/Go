package bitfield

import (
	"testing"
)

func TestBitfield(t *testing.T) {
	b := New(9)

	setInts := []int{1, 5, 8}
	checkInts := map[int]bool{
		1:   true,
		2:   false,
		3:   false,
		4:   false,
		5:   true,
		8:   true,
		100: false,
	}

	for _, i := range setInts {
		b.Set(i)
	}

	for k, v := range checkInts {
		ok := b.Contains(k)
		if ok != v {
			t.Fatalf(
				"\"%d\" is %ssupposed to be in the Bitfield.",
				k,
				map[bool]string{true: "", false: "NOT "}[v],
			)
		}

	}
}

func BenchmarkBitfield(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bf := New(b.N)

		for j := 0; j < b.N; j++ {
			bf.Set(j)
			bf.Contains(j)
		}
	}
}

func ExampleBitfield() {
	// creates a new bitfield with a minimum of 12 bits
	b := New(12)

	// set values
	b.Set(5)
	b.Set(1)

	// check if the bitfield contains values
	b.Contains(1) // true
	b.Contains(3) // false
	b.Contains(5) // true

	// reset the bitfield
	b.Reset()
	b.Contains(1) // false
	b.Contains(3) // false
	b.Contains(5) // false
}
