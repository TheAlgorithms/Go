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
	fmt.Println("")
	var b = t.Search(15)
	if b != nil {
		// fmt.Printf("%d %v\n", b.Val, b.isRed)
	} else {
		fmt.Println("Not found")
	}
	b = t.Search(50)

	if b != nil {
		// fmt.Printf("%d %v\n", b.Val, b.isRed)
	} else {
		fmt.Println("Not found")
	}

	t.Delete(5)
	fmt.Println("")
	t.PrintPreorder()
	fmt.Println("")
	t.PrintInorder()
	fmt.Println("")
	t.Delete(50)
	t.PrintPreorder()
	fmt.Println("")
	t.PrintInorder()
}
