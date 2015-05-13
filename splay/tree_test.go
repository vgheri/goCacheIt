package splay

import (
	"testing"
)

func createEmptyTree() Tree {
	fakeTree := new(Tree)
	fakeTree.Init()
	return *fakeTree
}

func createTreeWithRoot() Tree {
	fakeTree := new(Tree)
	fakeTree.Insert("myRoot", "{'test': 'value'}")
	return *fakeTree
}

func createPopulatedTree() Tree {
	fakeTree := new(Tree)
	fakeTree.Insert("Abc", "{'test': 'value_Abc'}")
	fakeTree.Insert("myRoot", "{'test': 'value'}")
	fakeTree.Insert("myRoot", "{'test': 'value'}")
	fakeTree.Insert("myRoot", "{'test': 'value'}")
	fakeTree.Insert("myRoot", "{'test': 'value'}")
	fakeTree.Insert("myRoot", "{'test': 'value'}")
	fakeTree.Insert("myRoot", "{'test': 'value'}")
	fakeTree.Insert("myRoot", "{'test': 'value'}")
	fakeTree.Insert("myRoot", "{'test': 'value'}")
	fakeTree.Insert("myRoot", "{'test': 'value'}")
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
	if fakeTree.root.value != fakeValue {
		t.Errorf("Expecting value: %s, found %s", fakeValue, fakeTree.root.value)
	}
}

func TestInsertNodeOnLeftOfRoot(t *testing.T) {
	fakeTree := createTreeWithRoot()
	fakeKey := "left"
	fakeValue := "{'test': 'value2'}"
	fakeTree.Insert(fakeKey, fakeValue)
	if fakeTree.root.right != nil {
		t.Error("Expecting root.right to be nil, it is not.")
	}
	node := fakeTree.root.left
	if node == nil {
		t.Error("Left node not added")
	}
	if node.key != fakeKey {
		t.Errorf("Expecting key: %s, found %s", fakeKey, node.key)
	}
	if node.value == nil {
		t.Errorf("Expecting value %s, found nil", fakeValue)
	}
	if node.value != fakeValue {
		t.Errorf("Expecting value: %s, found %s", fakeValue, node.value)
	}
}

func TestInsertNodeOnRightOfRoot(t *testing.T) {
	fakeTree := createTreeWithRoot()
	fakeKey := "right"
	fakeValue := "{'test': 'value2'}"
	fakeTree.Insert(fakeKey, fakeValue)
	if fakeTree.root.left != nil {
		t.Error("Expecting root.left to be nil, it is not.")
	}
	node := fakeTree.root.right
	if node == nil {
		t.Error("Right node not added")
	}
	if node.key != fakeKey {
		t.Errorf("Expecting key: %s, found %s", fakeKey, node.key)
	}
	if node.value == nil {
		t.Errorf("Expecting value %s, found nil", fakeValue)
	}
	if node.value != fakeValue {
		t.Errorf("Expecting value: %s, found %s", fakeValue, node.value)
	}
}
