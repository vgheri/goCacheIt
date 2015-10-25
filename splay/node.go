package splay

// Node is the type for tree elements
type Node struct {
	parent, left, right *Node
	key                 string
	Value               Any
}

func newNode(key string, value Any, parent *Node) *Node {
	return &Node{parent: parent, left: nil, right: nil, key: key, Value: value}
}

func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) isRoot() bool {
	return n.parent == nil
}
