package splay

import (
	"errors"
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

// Init initializes the Tree structure by setting the root node to nil
func (t *Tree) Init() {
	t.root = nil
}

// Insert a key-value couple into the tree
func (t *Tree) Insert(key string, value Any) error {
	var err error
	// if key is > key_max_length => error
	if len(key) > keyMaxLength {
		return errors.New("Key exceeding max length of 255 chars.")
	}
	if t.root == nil {
		t.root = &node{parent: nil, left: nil, right: nil, key: key, value: value}
		return err
	}
	current := t.root
	var parent *node
	var direction string
	for current != nil {
		if key == current.key {
			err = errors.New("Key already existing.")
			return err
		}
		parent = current
		if key < current.key {
			current = current.left
			direction = "left"
		} else {
			current = current.right
			direction = "right"
		}
	}
	current = &node{parent: parent, left: nil, right: nil, key: key, value: value}
	if direction == "left" {
		parent.left = current
	} else {
		parent.right = current
	}
	//splay
	return err
}
