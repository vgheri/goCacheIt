package splay

import (
	// "fmt"
	"testing"
)

func createFixedTree() *Tree {
	fakeTree := New(500)
	fakeTree.Insert("middle", "{'test': 'value_Abc'}")
	fakeTree.Insert("Amount", "{'test': 'value1'}")
	fakeTree.Insert("First", "{'test': 'value2'}")
	fakeTree.Insert("Delta", "{'test': 'value3'}")
	fakeTree.Insert("Geneve", "{'test': 'value4'}")
	fakeTree.Insert("netstat", "{'test': 'value5'}")
	fakeTree.Insert("nelly", "{'test': 'value6'}")
	fakeTree.Insert("nefertity", "{'test': 'value7'}")
	fakeTree.Insert("moriarty", "{'test': 'value5'}")
	fakeTree.Insert("polly", "{'test': 'value6'}")
	fakeTree.Insert("opportunity", "{'test': 'value7'}")
	fakeTree.Insert("sansa", "{'test': 'value7'}")
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
