// Cartesian tree is binary tree, heap-ordered,
// inorder traversal returns the original sequence,
// constructed by finding min values in a sequence
// Main property: parent node has smaller value than child node
// more info: https://en.wikipedia.org/wiki/Cartesian_tree

package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func BuildCartesianTree(sequenceOfNumbers []int, start int, end int) *Node {
	if start > end {
		return nil
	}
	index := findIndexOfMinValue(sequenceOfNumbers, start, end)
	root := createNewNode(sequenceOfNumbers[index])
	root.Left = BuildCartesianTree(sequenceOfNumbers, start, index-1)
	root.Right = BuildCartesianTree(sequenceOfNumbers, index+1, end)
	return root

}

func createNewNode(x int) *Node {
	newNode := &Node{x, nil, nil}
	return newNode
}

func findIndexOfMinValue(sequenceOfNumbers []int, start int, end int) int {
	minIndex := start
	for i := start; i <= end; i++ {
		if sequenceOfNumbers[i] < sequenceOfNumbers[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func inorderTraversal(root *Node) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	printNode(root)
	inorderTraversal(root.Right)
}

func printNode(n *Node) {
	fmt.Print("Value: ", n.Value)
	if n.Left != nil {
		fmt.Print(" Left node: ", n.Left.Value)
	}
	if n.Right != nil {
		fmt.Print(" Right node: ", n.Right.Value)
	}
	fmt.Println()
}

func main() {
	v := []int{9, 3, 7, 1, 8, 12, 10, 20, 15, 18, 5}
	root := BuildCartesianTree(v, 0, len(v)-1)
	printNode(root)
	inorderTraversal(root)
}
