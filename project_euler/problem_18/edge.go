package problem18

type Edge struct {
	ID        int
	NodeValue NodeValue
	NodeLeft  Node
	NodeRight Node
	Parent    Node
}

func (n *Edge) Value() NodeValue {
	return n.NodeValue
}

func (n *Edge) Left() Node {
	return n.NodeLeft
}

func (n *Edge) Right() Node {
	return n.NodeRight
}

func (n *Edge) Kind() string {
	return "edge"
}

func (n *Edge) CreateChild(value NodeValue, id int) Node {
	// When the left child is nil, it's a left edge
	if n.NodeLeft == nil {
		return &Edge{
			ID:        id,
			NodeValue: value,
			Parent:    n,
			NodeLeft:  nil,
			NodeRight: nil,
		}
	}

	// When the left child is a leaf, it's a right edge
	if n.NodeLeft.Kind() == "leaf" {
		return &Edge{
			ID:        id,
			NodeValue: value,
			Parent:    n,
			NodeLeft:  nil,
			NodeRight: nil,
		}
	}

	return &Leaf{
		ID:        id,
		NodeValue: value,
		Parent:    n,
		NodeLeft:  nil,
		NodeRight: nil,
	}
}

func (n *Edge) GetID() int {
	return n.ID
}

func (n *Edge) Insert(node Node) {
	// If Left is nil, always simply insert the node
	if n.NodeLeft == nil {
		node.SetParent(n)
		n.NodeLeft = node

		return
	}

	// If Right is nil, insert the node
	n.NodeRight = node

	// If the node to insert is an edge, set the parent
	if node.Kind() == "edge" {
		node.SetParent(n)

		return
	}

	// If the node to insert is a leaf, send it to the sibling right
	n.Parent.Right().Insert(node)
}

func (n *Edge) HasSpace() bool {
	return n.NodeLeft == nil || n.NodeRight == nil
}

func (n *Edge) LeftIsNil() bool {
	return n.NodeLeft == nil
}

func (n *Edge) RightIsNil() bool {
	return n.NodeRight == nil
}

func (n *Edge) SetParent(node Node) {
	n.Parent = node
}
