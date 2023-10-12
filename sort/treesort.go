//Implementation of tree sorting

package main

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value, Left: nil, Right: nil}
	}

	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else if value >= root.Value {
		root.Right = insert(root.Right, value)
	}

	return root
}

func inOrderTraversal(root *Node, sortedArr *[]int) {
	if root == nil {
		return
	}

	inOrderTraversal(root.Left, sortedArr)
	*sortedArr = append(*sortedArr, root.Value)
	inOrderTraversal(root.Right, sortedArr)
}

func treeSort(arr []int) []int {
	var root *Node

	for _, v := range arr {
		root = insert(root, v)
	}

	var sortedArr []int
	inOrderTraversal(root, &sortedArr)

	return sortedArr
}

func main() {
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter the elements:")

	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	sortedArr := treeSort(arr)

	fmt.Println("Sorted array:")
	fmt.Println(sortedArr)
}
