package splay

import (
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func createEmptyTree() *Tree {
	fakeTree := New()
	return fakeTree
}

func createTreeWithRoot(key string, value Any) *Tree {
	fakeTree := New()
	fakeTree.Insert(key, value)
	return fakeTree
}

func createDefaultTreeWithRoot() *Tree {
	return createTreeWithRoot("myRoot", "{'test': 'value'}")
}

func createPopulatedTree() *Tree {
	fakeTree := New()
	fakeTree.Insert(randSeq(5), "{'test': 'value_Abc'}")
	fakeTree.Insert(randSeq(5), "{'test': 'value1'}")
	fakeTree.Insert(randSeq(5), "{'test': 'value2'}")
	fakeTree.Insert(randSeq(5), "{'test': 'value3'}")
	fakeTree.Insert(randSeq(5), "{'test': 'value4'}")
	fakeTree.Insert(randSeq(5), "{'test': 'value5'}")
	fakeTree.Insert(randSeq(5), "{'test': 'value6'}")
	fakeTree.Insert(randSeq(5), "{'test': 'value7'}")
	return fakeTree
}

func TestInsertRootCorrectly(t *testing.T) {
	fakeKey := "myRoot"
	fakeValue := "{'test': 'value'}"
	fakeTree := createTreeWithRoot(fakeKey, fakeValue)
	if fakeTree.root == nil {
		t.Fatalf("Root node not added")
	}
	if fakeTree.root.key != "myRoot" {
		t.Fatalf("Expected key %s, found %s", fakeKey, fakeTree.root.key)
	}
	if fakeTree.root.value == nil {
		t.Fatalf("Expecting value not nil, found nil")
	}
	if fakeTree.root.value != fakeValue {
		t.Fatalf("Expected value %s, found %s", fakeValue, fakeTree.root.value)
	}
}

func TestInsertNodeOnLeftOfRoot(t *testing.T) {
	fakeTree := createDefaultTreeWithRoot()
	fakeKey := "left"
	fakeValue := "{'test': 'value2'}"
	fakeTree.Insert(fakeKey, fakeValue)
	if fakeTree.root.right != nil {
		t.Fatalf("Expecting root.right to be nil, it is not.")
	}
	node := fakeTree.root.left
	if node == nil {
		t.Fatalf("Left node not added")
		return
	}
	if node.key != fakeKey {
		t.Fatalf("Expecting key: %s, found %s", fakeKey, node.key)
	}
	if node.value == nil {
		t.Fatalf("Expecting value %s, found nil", fakeValue)
	}
	if node.value != fakeValue {
		t.Fatalf("Expecting value: %s, found %s", fakeValue, node.value)
	}
}

func TestInsertNodeOnRightOfRoot(t *testing.T) {
	fakeTree := createDefaultTreeWithRoot()
	fakeKey := "right"
	fakeValue := "{'test': 'value2'}"
	fakeTree.Insert(fakeKey, fakeValue)
	if fakeTree.root.left != nil {
		t.Fatalf("Expecting root.left to be nil, it is not.")
	}
	node := fakeTree.root.right
	if node == nil {
		t.Fatalf("Right node not added")
	}
	if node.key != fakeKey {
		t.Fatalf("Expecting key: %s, found %s", fakeKey, node.key)
	}
	if node.value == nil {
		t.Fatalf("Expecting value %s, found nil", fakeValue)
	}
	if node.value != fakeValue {
		t.Fatalf("Expecting value: %s, found %s", fakeValue, node.value)
	}
}

func TestCannotInsertDuplicateKey(t *testing.T) {
	fakeTree := createEmptyTree()
	err := fakeTree.Insert("a", "b")
	if err != nil {
		t.Fatal("It should have been able to insert root")
	}
	err2 := fakeTree.Insert("a", "c")
	if err2 == nil {
		t.Fatal("It should have raised an error")
	}
}

func TestCanInsertNodeAndCanGetItsValue(t *testing.T) {
	fakeTree := createPopulatedTree()
	var key string
	for {
		key = randSeq(5)
		if fakeTree.Get(key) == nil {
			break
		}
	}
	err := fakeTree.Insert(key, "{'test': 'abcdas'}")
	if err != nil {
		t.Fatal("It should have been able to insert node. Error: %s", err.Error())
	}
	if fakeTree.Get(key) == nil {
		t.Fatalf("It should have been able to find key %s", key)
	}
}
