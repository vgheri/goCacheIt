package splay

import (
	// "fmt"
	"testing"
	"time"
)

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

func TestMarkCorrectlyMarksNodeForDeletion(t *testing.T) {
	tree := createFixedTree()
	mark(tree)
	for _, n := range nodesToRemove {
		if n.key != "Delta" && n.key != "Geneve" && n.key != "moriarty" &&
			n.key != "opportunity" && n.key != "sansa" {
			t.Fatalf("Wrong nodes marked for deletion. Expected "+
				"Delta, Geneve, moriarty, opportunity and sansa. Found %s.", n.key)
		}
	}
}

func TestFreeMemoryRemovesMarkedNodesForDeletion(t *testing.T) {
	tree := createFixedTree()
	tree.freeMemory()
	tree.shouldContain("middle", t)
	tree.shouldContain("Amount", t)
	tree.shouldContain("First", t)
	tree.shouldContain("netstat", t)
	tree.shouldContain("nelly", t)
	tree.shouldContain("nefertity", t)
	tree.shouldContain("polly", t)
	tree.shouldNotContain("Delta", t)
	tree.shouldNotContain("Geneve", t)
	tree.shouldNotContain("moriarty", t)
	tree.shouldNotContain("opportunity", t)
	tree.shouldNotContain("sansa", t)
}

func TestRuningFreeMemoryMultipleTimes(t *testing.T) {
	tree := createFixedTree()
	tree.freeMemory()
	tree.shouldContain("middle", t)
	tree.shouldContain("Amount", t)
	tree.shouldContain("First", t)
	tree.shouldContain("netstat", t)
	tree.shouldContain("nelly", t)
	tree.shouldContain("nefertity", t)
	tree.shouldContain("polly", t)
	tree.shouldNotContain("Delta", t)
	tree.shouldNotContain("Geneve", t)
	tree.shouldNotContain("moriarty", t)
	tree.shouldNotContain("opportunity", t)
	tree.shouldNotContain("sansa", t)
	tree.freeMemory()
	tree.shouldContain("First", t)
	tree.shouldContain("netstat", t)
	tree.shouldContain("nelly", t)
	tree.shouldContain("nefertity", t)
	tree.shouldContain("polly", t)
	tree.shouldNotContain("middle", t)
	tree.shouldNotContain("Amount", t)
}

// func TestMemoryUsage(t *testing.T) {
// 	tree := createFixedTree()
// 	for {
// 		var key string
// 		for {
// 			key = randSeq(5)
// 			if node, _ := tree.Get(key); node == nil {
// 				break
// 			}
// 		}
// 		_, _ = tree.Insert(key, "{'test': 'asdasdadasasdasdadasasdasdadasasdasdadasasdasdadasasdasdadas'}")
// 	}
// 	close(tree.jobs)
// }

func TestExpiredKeyShouldBeRemovedFromTree(t *testing.T) {
	tree := createFixedTree()
	tree.Insert("somewhere", "{'test': 'value_Abc'}", 500*time.Millisecond)
	time.Sleep(1 * time.Second)
	if n, _ := tree.Get("somewhere"); n != nil {
		t.Fatal("Expired node should have been removed from the tree")
	}
}

func TestNotExpiredKeyShouldNotBeRemovedFromTree(t *testing.T) {
	tree := createFixedTree()
	tree.Insert("somewhere", "{'test': 'value_Abc'}", 5*time.Second)
	time.Sleep(1 * time.Second)
	if n, _ := tree.Get("somewhere"); n == nil {
		t.Fatal("Not expired node should not have been removed from the tree")
	}
}
