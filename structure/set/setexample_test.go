// package set implements a Set using a golang map.
// This implies that only the types that are accepted as valid map keys can be used as set elements.
// For instance, do not try to Add a slice, or the program will panic.

package set

import (
	"fmt"
)

func ExampleSet() {

	set := New(1, 2, 3)
	fmt.Println(set.Len()) // 3
	set.Add(3)
	fmt.Println(set.Len()) // 3
	set.Add(4)
	fmt.Println(set.Len()) // 4

	// output:
	// 3
	// 3
	// 4
}
