package disjointsetunion

import "fmt"
import "testing"

func Example() {

	var size int = 10
	var myDSU DSU
	myDSU.init(size)  // 10 numbers pointed to themselves. Size of individual elements initialized.
	myDSU.union(1, 2) // merging two sets
	myDSU.union(3, 2)
	myDSU.union(9, 3)
	parent := myDSU.find(3)                      // finding the parent
	fmt.Printf("%v\n", parent)                   // checking parent value of 3
	fmt.Printf("%v\n", myDSU.belongToSame(4, 9)) // Test to see if 2 elements belong to same set
	fmt.Printf("%v\n", myDSU.belongToSame(1, 3))
	myDSU.union(4, 9)
	fmt.Printf("%v\n", myDSU.belongToSame(4, 9))
}

func BenchmarkDSU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var myDSU DSU
		size := 100
		myDSU.init(size)
		myDSU.union(4, 48)
		myDSU.union(30, 6)
		parent := myDSU.find(48)
		fmt.Printf("%v\n", parent)                     // 4
		fmt.Printf("%v\n", myDSU.belongToSame(6, 30))  // true
		fmt.Printf("%v\n", myDSU.belongToSame(8, 50))  //false
		fmt.Printf("%v\n", myDSU.belongToSame(90, 30)) // false
	}
}
