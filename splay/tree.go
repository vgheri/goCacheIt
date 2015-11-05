package splay

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// expressed in MegaBytes
var maxMemory uint64

// Global variables
var cacheEvictionTicker *time.Ticker

// Any defines a generic type accepted by the Tree as a value
type Any interface{}

// Tree is the basic type for the splay package
type Tree struct {
	root *Node
	jobs chan *job
}

type job struct {
	command string
	node    *Node
	done    chan *Node
}

// New initializes the Tree structure by setting the root node to nil
// maxAmountOfMemory is expressed in MegaBytes
// TODO New should take an option object as input
func New(maxAmountOfMemory uint64) *Tree {
	t := new(Tree)
	t.root = nil
	t.jobs = make(chan *job)
	cacheEvictionTicker = time.NewTicker(memoryCheckFrequency)
	maxMemory = maxAmountOfMemory
	go t.workerLoop()
	return t
}

// worker loop
func (t *Tree) workerLoop() {
	for {
		select {
		case job, more := <-t.jobs:
			if more {
				if job != nil {
					var node *Node
					switch job.command {
					case commandInsertNode:
						node = t.insertNode(job.node, t.root, nil)
					case commandGetNode:
						node = t.getNode(job.node.key, t.root)
						splay(t, node)
					case commandRemoveNode:
						node = t.removeNode(job.node)
					}
					if job.done != nil {
						job.done <- node
					}
				}
			} else {
				cacheEvictionTicker.Stop()
				return
			}
		case <-cacheEvictionTicker.C:
			t.purgeNodes()
		}
	}
}

// Insert adds a key-value couple into the tree
func (t *Tree) Insert(key string, value Any, duration time.Duration) (*Node, error) {
	if !keyIsValid(key) {
		return nil, errors.New("Invalid key.")
	}
	if n, _ := t.Get(key); n != nil {
		return nil, errors.New("Key already exists.")
	}

	// create a job and send it to the jobs channel
	done := make(chan *Node)
	job := &job{
		command: commandInsertNode,
		node:    newNode(key, value, nil, duration),
		done:    done,
	}
	t.jobs <- job
	// wait for the operation to complete
	node := <-done
	return node, nil
}

// Get retrieves a node by key. Nil if the key doesn't exist
func (t *Tree) Get(key string) (*Node, error) {
	if !keyIsValid(key) {
		return nil, errors.New("Invalid key.")
	}

	// create a job and send it to the jobs channel
	done := make(chan *Node)
	job := &job{
		command: commandGetNode,
		node:    newNode(key, nil, nil, 1*time.Hour),
		done:    done,
	}
	t.jobs <- job
	// wait for the operation to complete
	node := <-done

	return node, nil
}

// Remove deletes the node with the desired key from the tree.
// Error if the key does not exist
func (t *Tree) Remove(key string) (*Node, error) {
	if !keyIsValid(key) {
		return nil, errors.New("Invalid key.")
	}
	var node *Node
	if node, _ = t.Get(key); node == nil {
		return nil, errors.New("Key does not exist.")
	}

	// create a job and send it to the jobs channel
	done := make(chan *Node)
	job := &job{
		command: commandRemoveNode,
		node:    node,
		done:    done,
	}
	t.jobs <- job
	// wait for the operation to complete
	_ = <-done

	return node, nil
}

/*** Support functions ***/
func (t *Tree) insertNode(node *Node, current, parent *Node) *Node {
	if current == nil {
		t.setRoot(node)
		return node
	}
	switch compare(node.key, current.key) {
	case -1:
		if current.left == nil {
			node.parent = current
			current.left = node
			return node
		}
		return t.insertNode(node, current.left, current)
	case 1:
		if current.right == nil {
			node.parent = current
			current.right = node
			return node
		}
		return t.insertNode(node, current.right, current)
	}
	return nil
}

func (t *Tree) getNode(key string, node *Node) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		return t.getNode(key, node.left)
	} else if key > node.key {
		return t.getNode(key, node.right)
	} else { // hit!
		return node
	}
}

func (t *Tree) removeNode(node *Node) *Node {
	if node.left == nil {
		replace(t, node, node.right)
	} else if node.right == nil {
		replace(t, node, node.left)
	} else {
		minimum := subtreeMinimum(node.right)
		if minimum.parent != node {
			replace(t, minimum, minimum.right)
			minimum.right = node.right
			minimum.right.parent = minimum
		}
		replace(t, node, minimum)
		minimum.left = node.left
		minimum.left.parent = minimum
	}
	return node
}

func keyIsValid(key string) bool {
	if len(key) > keyMaxLength {
		return false
	}
	return true
}

func (t *Tree) setRoot(node *Node) {
	node.parent = nil
	t.root = node
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

func printNode(n *Node, depth int) {
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
	fmt.Printf("%s%s[%s][Expiration %s]\n", strings.Repeat("-", 2*depth), side, n.key, n.expirationDate.Format("Mon Jan _2 15:04:05 2006"))
	printNode(n.left, depth+1)
	printNode(n.right, depth+1)
}

///////////////////////////////////////////////
/// Code block taken from Wikipedia         ///
/// http://en.wikipedia.org/wiki/Splay_tree ///
///////////////////////////////////////////////
func leftRotate(t *Tree, x *Node) {
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

func rightRotate(t *Tree, x *Node) {
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

func splay(t *Tree, x *Node) {
	if x == nil {
		return
	}
	if parent := x.parent; parent != nil {
		if parent.parent == nil {
			if parent.left == x {
				rightRotate(t, parent)
			} else {
				leftRotate(t, parent)
			}
		} else {
			if parent.left == x && parent.parent.left == parent {
				rightRotate(t, parent.parent)
				rightRotate(t, parent)
			} else if parent.right == x && parent.parent.right == parent {
				leftRotate(t, parent.parent)
				leftRotate(t, parent)
			} else if parent.left == x && parent.parent.right == parent {
				rightRotate(t, parent)
				leftRotate(t, parent)
			} else {
				leftRotate(t, parent)
				rightRotate(t, parent)
			}
		}
	}
}

func replace(t *Tree, u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func subtreeMinimum(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}
