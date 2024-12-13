package problem18

type Leaf struct {
	ID        int
	NodeValue NodeValue
	NodeLeft  *Leaf
	NodeRight *Leaf
	Parent    Node
}

func (n *Leaf) Value() NodeValue {
	return n.NodeValue
}

func (n *Leaf) Left() Node {
	if n.NodeLeft != nil {
		n.NodeLeft.Parent = n // Leaf is the parent of its left child always
	}

	return n.NodeLeft
}

func (n *Leaf) Right() Node {
	return n.NodeRight
}

func (n *Leaf) Kind() string {
	return "leaf"
}

func (n *Leaf) CreateChild(value NodeValue, id int) Node {
	// Leafs only have leaf children
	return &Leaf{
		ID:        id,
		NodeValue: value,
		Parent:    n,
		NodeLeft:  nil,
		NodeRight: nil,
	}
}

func (n *Leaf) GetID() int {
	return n.ID
}

func (n *Leaf) Insert(node Node) {
	// If Left is nil, always simply insert the node
	if n.NodeLeft == nil {
		node.SetParent(n)
		n.NodeLeft = node.(*Leaf)

		return
	}

	// If Right is nil, insert the node
	n.NodeRight = node.(*Leaf)
	// Send it to the sibling right
	n.Parent.Right().Insert(node)
}

func (n *Leaf) HasSpace() bool {
	return n.NodeLeft == nil || n.NodeRight == nil
}

func (n *Leaf) LeftIsNil() bool {
	return n.NodeLeft == nil
}

func (n *Leaf) RightIsNil() bool {
	return n.NodeRight == nil
}

func (n *Leaf) SetParent(node Node) {
	n.Parent = node
}
