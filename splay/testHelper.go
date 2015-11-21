package splay

import (
	"testing"
	"time"
)

var defaultDuration = 1 * time.Hour

func createFixedTree() *Tree {
	fakeTree := New(500)
	insertNodeWithDefaultDuration(fakeTree, "middle", "{'test': 'value_Abc'}")
	insertNodeWithDefaultDuration(fakeTree, "Amount", "{'test': 'value1'}")
	insertNodeWithDefaultDuration(fakeTree, "First", "{'test': 'value2'}")
	insertNodeWithDefaultDuration(fakeTree, "Delta", "{'test': 'value3'}")
	insertNodeWithDefaultDuration(fakeTree, "Geneve", "{'test': 'value4'}")
	insertNodeWithDefaultDuration(fakeTree, "netstat", "{'test': 'value5'}")
	insertNodeWithDefaultDuration(fakeTree, "nelly", "{'test': 'value6'}")
	insertNodeWithDefaultDuration(fakeTree, "nefertity", "{'test': 'value7'}")
	insertNodeWithDefaultDuration(fakeTree, "moriarty", "{'test': 'value5'}")
	insertNodeWithDefaultDuration(fakeTree, "polly", "{'test': 'value6'}")
	insertNodeWithDefaultDuration(fakeTree, "opportunity", "{'test': 'value7'}")
	insertNodeWithDefaultDuration(fakeTree, "sansa", "{'test': 'value7'}")
	return fakeTree
}

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
