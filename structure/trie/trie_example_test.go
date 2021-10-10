package trie

import "fmt"

func ExampleNode() {
	// creates a new node
	node := NewNode()

	// adds words
	node.Insert("nikola")
	node.Insert("tesla")

	// check size and capacity
	fmt.Println(node.Size())     // 2 words
	fmt.Println(node.Capacity()) // 12 nodes

	// finds words
	fmt.Println(node.Find("thomas")) // false
	fmt.Println(node.Find("edison")) // false
	fmt.Println(node.Find("nikola")) // true

	// remove a word, and check it is gone
	node.Remove("tesla")
	fmt.Println(node.Find("tesla")) // false

	// size and capacity have changed
	fmt.Println(node.Size())     // 1 word left
	fmt.Println(node.Capacity()) // 12 nodes remaining

	// output:
	// 2
	// 12
	// false
	// false
	// true
	// false
	// 1
	// 12
}
