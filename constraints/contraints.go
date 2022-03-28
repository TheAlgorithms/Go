// Package constraints has some useful generic type constraints defined which is very similar to
// [golang.org/x/exp/constraints](https://pkg.go.dev/golang.org/x/exp/constraints) package.
// We refrained from using that until it gets placed into the standard library - currently
// there are some questions regarding this package [ref](https://github.com/golang/go/issues/50792).
package constraints

// Signed is a generic type constraint for all signed integers.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a generic type constraint for all unsigned integers.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Integer is a generic type constraint for all integers (signed and unsigned.)
type Integer interface {
	Signed | Unsigned
}

// Float is a generic type constraint for all floating point types.
type Float interface {
	~float32 | ~float64
}

// Number is a generic type constraint for all numeric types in Go except Complex types.
type Number interface {
	Integer | Float
}

// Ordered is a generic type constraint for all ordered data types in Go.
// Loosely speaking, in mathematics a field is an ordered field if there is a "total
// order" (a binary relation - in this case `<` symbol) such that we will always have
// if a < b then a + c < b + c and if 0 < a, 0 < b then 0 < a.b
// The idea in Go is quite similar, though only limited to Go standard types
// not user defined types.
type Ordered interface {
	Integer | ~string | Float
}
