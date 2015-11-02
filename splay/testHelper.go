package splay

import (
	"testing"
	"time"
)

var defaultDuration = 1 * time.Hour

func insertNodeWithDefaultDuration(tree *Tree, key string, value Any) (*Node, error) {
	return tree.Insert(key, value, defaultDuration)
}

func (tree *Tree) shouldContain(key string, t *testing.T) {
	if n, _ := tree.Get(key); n == nil {
		t.Fatalf("Expected to find node %s. Found none.", key)
	}
}

func (tree *Tree) shouldNotContain(key string, t *testing.T) {
	if n, _ := tree.Get(key); n != nil {
		t.Fatalf("It should not have found node %s. Found it instead.", key)
	}
}
