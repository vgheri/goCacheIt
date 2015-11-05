package splay

import (
	"time"
)

// Node is the type for tree elements
type Node struct {
	parent, left, right *Node
	key                 string
	Value               Any
	expirationDate      time.Time
}

func newNode(key string, value Any, parent *Node, duration time.Duration) *Node {
	return &Node{
		parent:         parent,
		left:           nil,
		right:          nil,
		key:            key,
		Value:          value,
		expirationDate: time.Now().Add(duration),
	}
}

func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) isRoot() bool {
	return n.parent == nil
}

func (n *Node) isExpired() bool {
	return n.expirationDate.Before(time.Now())
}
