package splay

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

const keyMaxLength int = 255

// Any defines a generic type accepted by the Tree as a value
type Any interface{}

type node struct {
	parent, left, right *node
	key                 string
	value               Any
	lock                *sync.Mutex
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
	if t.Get(key) != nil {
		return errors.New("Key already exists.")
	}
	node := insertNode(key, value, t.root, nil, t)

	splay(t, node)
	return nil
}

// Get retrieves a value by key. Nil if the key doesn't exist
func (t *Tree) Get(key string) Any {
	node := getNode(key, t.root)
	if node == nil {
		return nil
	}
	splay(t, node)
	return node.value
}

/*** Support functions ***/
func newNode(key string, value Any, parent *node) *node {
	return &node{parent: parent, left: nil, right: nil, key: key, value: value, lock: &sync.Mutex{}}
}

func insertNode(key string, value Any, current, parent *node, t *Tree) *node {
	if current == nil {
		current = newNode(key, value, parent)
		t.root = current
		return current
	}
	switch compare(key, current.key) {
	case -1:
		if current.left == nil {
			current.left = newNode(key, value, current)
			return current.left
		}
		return insertNode(key, value, current.left, current, t)
	case 1:
		if current.right == nil {
			current.right = newNode(key, value, current)
			return current.right
		}
		return insertNode(key, value, current.right, current, t)
	}
	return nil
}

func getNode(key string, node *node) *node {
	if node == nil {
		return nil
	}
	if key < node.key {
		return getNode(key, node.left)
	} else if key > node.key {
		return getNode(key, node.right)
	} else { // hit!
		return node
	}
}

func keyIsValid(key string) bool {
	if len(key) > keyMaxLength {
		return false
	}
	return true
}

func (t *Tree) setRoot(key string, value Any) {
	t.root = newNode(key, value, nil)
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

func (t *Tree) print() {
	if t == nil || t.root == nil {
		fmt.Println("Empty tree")
		return
	}
	printNode(t.root, 0)
}

func printNode(n *node, depth int) {
	if n == nil {
		return
	}
	side := ""
	if n.parent != nil {
		if n.parent.right == n {
			side = "(R)"
		} else {
			side = "(L)"
		}
	}
	fmt.Printf("%s%s[%s]\n", strings.Repeat("-", 2*depth), side, n.key)
	printNode(n.left, depth+1)
	printNode(n.right, depth+1)
}

///////////////////////////////////////////////
/// Code block taken from Wikipedia         ///
/// http://en.wikipedia.org/wiki/Splay_tree ///
///////////////////////////////////////////////
func leftRotate(t *Tree, x *node) {
	y := x.right
	if y != nil {
		x.right = y.left
		if y.left != nil {
			y.left.parent = x
		}
		y.parent = x.parent
	}

	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	if y != nil {
		y.left = x
	}
	x.parent = y
}

func rightRotate(t *Tree, x *node) {
	y := x.left
	if y != nil {
		x.left = y.right
		if y.right != nil {
			y.right.parent = x
		}
		y.parent = x.parent
	}
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	if y != nil {
		y.right = x
	}
	x.parent = y
}

func splay(t *Tree, x *node) {
	if x == nil {
		return
	}
	if x.parent != nil {
		if x.parent.parent == nil {
			if x.parent.left == x {
				rightRotate(t, x.parent)
			} else {
				leftRotate(t, x.parent)
			}
		} else if x.parent.left == x && x.parent.parent.left == x.parent {
			rightRotate(t, x.parent.parent)
			rightRotate(t, x.parent)
		} else if x.parent.right == x && x.parent.parent.right == x.parent {
			leftRotate(t, x.parent.parent)
			leftRotate(t, x.parent)
		} else if x.parent.left == x && x.parent.parent.right == x.parent {
			rightRotate(t, x.parent)
			leftRotate(t, x.parent)
		} else {
			leftRotate(t, x.parent)
			rightRotate(t, x.parent)
		}
	}
}
