package splay

import (
	"testing"
)

func createEmptyTree() Tree {
	fakeTree := new(Tree)
	fakeTree.Init()
	return *fakeTree
}

func TestInsertRootCorrectly(t *testing.T) {
	fakeTree := createEmptyTree()
	fakeKey := "myRoot"
	fakeValue := "{'test': 'value'}"
	fakeTree.Insert(fakeKey, fakeValue)
	if fakeTree.root == nil {
		t.Error("Root node not added")
	}
	if fakeTree.root.key != "myRoot" {
		t.Errorf("Expecting key: %s, found %s", fakeKey, fakeTree.root.key)
	}
	if fakeTree.root.value == nil {
		t.Error("Expecting value not nil, found nil")
	}
	if fakeTree.root.value != "{'test': 'value'}" {
		t.Errorf("Expecting value: %s, found %s", fakeValue, fakeTree.root.value)
	}
}
