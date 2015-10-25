package splay

import (
	"runtime"
)

var nodesToRemove []*Node

func shouldFreeMemory() bool {
	shouldFreeMemory := false
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	usedMemory := stats.HeapSys / 1000000
	if usedMemory > maxMemory*uint64(memoryUsageThreshold) {
		shouldFreeMemory = true
	}
	return shouldFreeMemory
}

func freeMemory(t *Tree) {
	// mark
	mark(t)
	// sweep
	sweep(t)
}

func mark(t *Tree) {
	markNode(t.root)
}

func markNode(node *Node) {
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
		removeNode(node, t)
	}
}
