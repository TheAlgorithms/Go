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
