package splay

import (
	"runtime"
)

var nodesToRemove []*Node

func shouldFreeMemory() bool {
	shouldFreeMemory := false
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	usedMemory := stats.HeapAlloc
	threshold := ((maxMemory * 1000000) * uint64(memoryUsageThreshold) / 100)
	if usedMemory > threshold {
		shouldFreeMemory = true
	}
	return shouldFreeMemory
}

func (t *Tree) freeMemory() {
	// TODO improve initial capacity to a possibly meaningful number
	nodesToRemove = make([]*Node, 0)
	// mark nodes for deletion
	mark(t)
	// remove marked nodes from the tree
	sweep(t)
}

func mark(t *Tree) {
	if t.root.left != nil || t.root.right != nil {
		markNode(t.root)
	}
}

func markNode(node *Node) {
	if node == nil {
		panic("In markNode: node should not be nil")
	}
	if node.left == nil && node.right == nil {
		nodesToRemove = append(nodesToRemove, node)
		return
	}
	if node.left != nil {
		markNode(node.left)
	}
	if node.right != nil {
		markNode(node.right)
	}
}

func sweep(t *Tree) {
	for _, node := range nodesToRemove {
		t.removeNode(node)
	}
}
