package problem18

type Root struct {
	ID        int
	NodeValue NodeValue
	NodeLeft  *Edge
	NodeRight *Edge
}

func (n *Root) Value() NodeValue {
	return n.NodeValue
}

func (n *Root) Left() Node {
	return n.NodeLeft
}

func (n *Root) Right() Node {
	return n.NodeRight
}

func (n *Root) Kind() string {
	return "root"
}

func (n *Root) CreateChild(value NodeValue, id int) Node {
	return &Edge{
		ID:        id,
		NodeValue: value,
		Parent:    n,
		NodeLeft:  nil,
		NodeRight: nil,
	}
}

func (n *Root) GetID() int {
	return n.ID
}

func (n *Root) Insert(node Node) {
	if n.NodeLeft == nil {
		n.NodeLeft = node.(*Edge)
	} else {
		n.NodeRight = node.(*Edge)
	}
}

func (n *Root) HasSpace() bool {
	return n.NodeLeft == nil || n.NodeRight == nil
}

func (n *Root) LeftIsNil() bool {
	return n.NodeLeft == nil
}

func (n *Root) RightIsNil() bool {
	return n.NodeRight == nil
}

func (n *Root) SetParent(node Node) {
	panic("Root node cannot have a parent")
}
