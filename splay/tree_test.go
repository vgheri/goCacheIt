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
	fakeTree := New(1)
	return fakeTree
}

func createTreeWithRoot(key string, value Any) *Tree {
	fakeTree := New(1)
	insertNodeWithDefaultDuration(fakeTree, key, value)
	return fakeTree
}

func createDefaultTreeWithRoot() *Tree {
	return createTreeWithRoot("myRoot", "{'test': 'value'}")
}

func createPopulatedTree() *Tree {

	fakeTree := New(1)
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'value_Abc'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'value1'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'value2'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'value3'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'value4'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'value5'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'value6'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'value7'}")
	return fakeTree
}

func TestInsertRootCorrectly(t *testing.T) {
	fakeKey := "myRoot"
	fakeValue := "{'test': 'value'}"
	fakeTree := createTreeWithRoot(fakeKey, fakeValue)
	close(fakeTree.jobs)
	if fakeTree.root == nil {
		t.Fatalf("Root node not added")
	}
	if fakeTree.root.key != "myRoot" {
		t.Fatalf("Expected key %s, found %s", fakeKey, fakeTree.root.key)
	}
	if fakeTree.root.Value == nil {
		t.Fatalf("Expecting value not nil, found nil")
	}
	if fakeTree.root.Value != fakeValue {
		t.Fatalf("Expected value %s, found %s", fakeValue, fakeTree.root.Value)
	}
}

func TestCannotInsertDuplicateKey(t *testing.T) {
	fakeTree := createEmptyTree()
	_, err := insertNodeWithDefaultDuration(fakeTree, "a", "b")
	if err != nil {
		t.Fatal("It should have been able to insert node")
	}
	_, err2 := insertNodeWithDefaultDuration(fakeTree, "a", "c")
	if err2 == nil {
		t.Fatal("It should have raised an error")
	}
	close(fakeTree.jobs)
}

func TestCanInsertNodeAndCanGetItsValue(t *testing.T) {
	fakeTree := createPopulatedTree()
	var key string
	for {
		key = randSeq(5)
		if node, _ := fakeTree.Get(key); node == nil {
			break
		}
	}
	insertNodeWithDefaultDuration(fakeTree, key, "{'test': 'abcdas'}")
	if node, _ := fakeTree.Get(key); node == nil {
		t.Fatalf("It should have been able to find key %s", key)
	}
	close(fakeTree.jobs)
}

func TestGetNonExistentNodeReturnsNil(t *testing.T) {
	fakeTree := createEmptyTree()
	insertNodeWithDefaultDuration(fakeTree, "a", "b")
	insertNodeWithDefaultDuration(fakeTree, "S", "c")
	insertNodeWithDefaultDuration(fakeTree, "f", "e")
	value, _ := fakeTree.Get("T")
	if value != nil {
		t.Fatal("Getting a non existent key should have returned nil")
	}
}

func TestGetSameKeyShouldEventuallyMoveNodeToRoot(t *testing.T) {
	fakeTree := createPopulatedTree()
	var key string
	for {
		key = randSeq(5)
		if node, _ := fakeTree.Get(key); node == nil {
			break
		}
	}
	insertNodeWithDefaultDuration(fakeTree, key, "{'test': 'abcdas'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'abcdas'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'abcdas'}")
	insertNodeWithDefaultDuration(fakeTree, randSeq(5), "{'test': 'abcdas'}")

	maxIterations := 10
	iterations := 0
	for fakeTree.root.key != key {
		if node, _ := fakeTree.Get(key); node == nil {
			t.Fatalf("It should have been able to find key %s", key)
		}
		iterations++
		if iterations == maxIterations {
			t.Fatal("It should have already moved the node to the root")
		}
	}
}

func TestRemoveReturnErrorIfKeyDoesntExist(t *testing.T) {
	fakeTree := createTreeWithRoot("test", "test 2")
	if _, err := fakeTree.Remove("notexistent"); err == nil {
		t.Fatalf("It should have thrown an error on removing non existent key")
	}
}

var result error

func BenchmarkInsert(b *testing.B) {
	fakeTree := createPopulatedTree()
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var key string
		for {
			key = randSeq(5)
			if node, _ := fakeTree.Get(key); node == nil {
				break
			}
		}
		_, err = insertNodeWithDefaultDuration(fakeTree, key, "{'test': 'abcdas'}")
	}
	result = err
	close(fakeTree.jobs)
}

func BenchmarkParallelInsert(b *testing.B) {
	fakeTree := createPopulatedTree()
	var err error
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var key string
			for {
				key = randSeq(5)
				if node, _ := fakeTree.Get(key); node == nil {
					break
				}
			}
			_, err = insertNodeWithDefaultDuration(fakeTree, key, "{'test': 'abcdas'}")
		}
	})
	result = err
	close(fakeTree.jobs)
}
