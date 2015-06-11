package splay

import (
	// "fmt"
	"math/rand"
	"testing"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		ran := r.Intn(len(letters))
		b[i] = letters[ran]
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
	fakeKey := "testLeft"
	fakeValue := "{'testLeft': 'value2'}"
	fakeTree.Insert(fakeKey, fakeValue)
	if fakeTree.root.key != "testLeft" {
		t.Fatalf("Expecting new inserted node to be root, it is not.")
	}
	if fakeTree.root.left.key != "myRoot" {
		t.Fatal("Expected former root node to be new root's left node, but it wasn't.")
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
	fakeTree.Insert(key, "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	if fakeTree.Get(key) == nil {
		t.Fatalf("It should have been able to find key %s", key)
	}
}

func TestGetNonExistentNodeReturnsNil(t *testing.T) {
	fakeTree := createEmptyTree()
	fakeTree.Insert("a", "b")
	fakeTree.Insert("S", "c")
	fakeTree.Insert("f", "e")
	value := fakeTree.Get("T")
	if value != nil {
		t.Fatal("Getting a non existent key should have returned nil")
	}
}

func TestGetSameKeyShouldEventuallyMoveNodeToRoot(t *testing.T) {
	fakeTree := createPopulatedTree()
	var key string
	for {
		key = randSeq(5)
		if fakeTree.Get(key) == nil {
			break
		}
	}
	fakeTree.Insert(key, "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	maxIterations := 10
	iterations := 0
	for fakeTree.root.key != key {
		if fakeTree.Get(key) == nil {
			t.Fatalf("It should have been able to find key %s", key)
		}
		iterations++
		if iterations == maxIterations {
			t.Fatal("It should have already moved the node to the root")
		}
	}
}

func TestCannotModifyTreeIfRootIsLocked(t *testing.T) {
	fakeTree := createDefaultTreeWithRoot()
	fakeTree.root.lock.Lock()
	go func() {
		fakeTree.root.lock.Lock()
		fakeTree.root.lock.Unlock()
	}()
	time.Sleep(time.Second)
	fakeTree.root.lock.Unlock()
}

func TestRemoveSuccessfullyRemovesNodeFromTree(t *testing.T) {
	fakeTree := createPopulatedTree()
	var key string
	for {
		key = randSeq(5)
		if fakeTree.Get(key) == nil {
			break
		}
	}
	fakeTree.Insert(key, "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	fakeTree.Insert(randSeq(5), "{'test': 'abcdas'}")
	if err := fakeTree.Remove(key); err != nil {
		t.Fatalf("It should have successfully removed key %s, got an error instead", key)
	}
	if n := fakeTree.Get(key); n != nil {
		t.Fatalf("Found key %s, expected to be removed from the tree", key)
	}
}

func TestRemoveReturnErrorIfKeyDoesntExist(t *testing.T) {
	fakeTree := createTreeWithRoot("test", "test 2")
	if err := fakeTree.Remove("notexistent"); err == nil {
		t.Fatalf("It should have thrown an error on removing non existent key")
	}
}
