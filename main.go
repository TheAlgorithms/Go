package main

import (
	"fmt"

	"github.com/TheAlgorithms/Go/structure/redblacktree"
)

func main() {
	// fmt.Println("Hey")
	// var a = redblacktree.Node{}
	var t = redblacktree.RBTree{}
	var vals = []int{8, 18, 5, 15, 17, 25, 40, 80}
	for _, val := range vals {
		t.Insert(val)
		// t.Print()
	}
	t.PrintPreorder()
	fmt.Println("")
	t.PrintInorder()
}
