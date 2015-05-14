package splay

import (
	"errors"
	// "fmt"
)

const keyMaxLength int = 255

// Any defines a generic type accepted by the Tree as a value
type Any interface{}

type node struct {
	parent, left, right *node
	key                 string
	value               Any
}

// Tree is the basic type for the splay package
type Tree struct {
	root *node
}

// New initializes the Tree structure by setting the root node to nil
func New() *Tree {
	t := new(Tree)
	t.root = nil
	return t
}

// Insert a key-value couple into the tree
func (t *Tree) Insert(key string, value Any) error {
	if !keyIsValid(key) {
		return errors.New("Invalid key.")
	}
	// if t.root == nil {
	// 	t.setRoot(key, value)
	// 	return nil
	// }
	if t.Get(key) != nil {
		return errors.New("Key already exists.")
	}
	insertNode(key, value, t.root, nil, t)

	//splay
	return nil
}

// Get retrieves a value by key. Nil if the key doesn't exist
func (t *Tree) Get(key string) Any {
	return getNode(key, t.root)
}

/*** Support functions ***/
func insertNode(key string, value Any, current, parent *node, t *Tree) *node {
	if current == nil {
		current = &node{parent: parent, left: nil, right: nil, key: key, value: value}
		t.root = current
		return current
	}
	switch compare(key, current.key) {
	case -1:
		if current.left == nil {
			current.left = &node{parent: current, left: nil, right: nil, key: key, value: value}
			return current.left
		}
		return insertNode(key, value, current.left, current, t)
	case 1:
		if current.right == nil {
			current.right = &node{parent: current, left: nil, right: nil, key: key, value: value}
			return current.right
		}
		return insertNode(key, value, current.right, current, t)
	}
	return nil
}

func getNode(key string, node *node) Any {
	if node == nil {
		return nil
	}
	if key < node.key {
		return getNode(key, node.left)
	} else if key > node.key {
		return getNode(key, node.right)
	} else { // hit!
		// splay
		return node.value
	}
}

func keyIsValid(key string) bool {
	if len(key) > keyMaxLength {
		return false
	}
	return true
}

func (t *Tree) setRoot(key string, value Any) {
	t.root = &node{parent: nil, left: nil, right: nil, key: key, value: value}
}

func compare(a, b string) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}
