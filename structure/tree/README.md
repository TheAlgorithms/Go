## Documentation for `tree` package:

1. `NewRBTree`:  Create a new Red-Black Tree
`tree := NewRBTree[int]()` without any parameters
2. `NewAVLTree`:  Create a new AVL Tree
`tree := NewAVLTree[int]()` without any parameters
3. `NewBSTree`:  Create a new BST
`tree := NewBSTree[int]()` without any parameters
4. `Insert`:  Insert a chain of keys into the tree.  Implemented by **BST, AVL and Red-Black Tree**.
```go
tree := NewBSTree[int]()  // or tree := NewRBTree[int]() or tree := NewAVLTree[int]()
nums := []int{1,2,3,4}
tree.Insert(nums...)
tree.Insert(10)
tree.Insert(10, 20)
// The above three ways make sense.
```
6. `InOrder`, `PreOrder` , `PostOrder` and `LevelOrder` :  Traverse the tree via four different ways. Implemented by **BST, AVL and Red-Black Tree**.
```go
tree.Insert(nums...)
ret := tree.PreOrder() 
ret = tree.InOrder()
ret = tree.PostOrder()
ret = tree.LevelOrder()
// ret is a slice of ints
```
7. `Delete`:  Delete a node of Tree.  Implemented by **BST, AVL and Red-Black Tree**.
```go
tree := NewAVLTree[int]()
tree.Insert(5)
tree.Insert(4)
tree.Insert(3)
tree.Insert(2)

tree.Delete(5)
```
8. `Max` and `Min` :  Return the max or min value of the tree. Implemented by **BST, AVL and Red-Black Tree**.
```go
key, ok := tree.Max() // if the tree is empty, ok is false, otherwise it's true.
```
10. `Has` :  Determine the tree has the node of key. Implemented by **BST, AVL and Red-Black Tree**.
```go
tree := NewAVLTree[int]()
tree.Insert(5)
tree.Insert(4)
tree.Insert(3)
tree.Insert(2) 
have := tree.Has(5) 
```
12. `Get`: Get the node of key. Implemented by **BST, AVL and Red-Black Tree**.
```go
tree := NewAVLTree[int]()
tree.Insert(5)
tree.Insert(4)
tree.Insert(3)
tree.Insert(2) 
n, ok := tree.Get(15) // if 15 isn't in the tree, ok is false, otherwise it's true.
```
12. `Depth`: Returns the depth of the tree. Implemented by **BST, AVL and Red-Black Tree**.

14. `Predecessor`:  Return the predecessor of the node of key. Only Implemented by **Red-Black Tree**.
```go
tree := NewRBTree[int]()
nums := []int{10, 8, 88, 888, 4, -1, 100}
tree.Insert(nums...)
ret, ok := tree.Predecessor(100) // ret = 88, ok = true
// If there is no predecessor, ok = false
```
15. `Successor`:  Return the successor of the node of key. Only Implemented by **Red-Black Tree**.
```go
tree := NewRBTree[int]()
nums := []int{10, 8, 88, 888, 4, -1, 100}
tree.Insert(nums...)
ret, ok := tree.Successor(100) // ret = 888, ok = true
// If there is no successor, ok = false
```