package splay

import (
	"testing"
	"time"
)

func TestMarkCorrectlyMarksNodeForDeletion(t *testing.T) {
	tree := createFixedTree()
	shouldDeleteLeaves = true
	mark(tree)
	for _, n := range nodesToRemove {
		if n.key != "Delta" && n.key != "Geneve" && n.key != "moriarty" &&
			n.key != "opportunity" && n.key != "sansa" {
			t.Fatalf("Wrong nodes marked for deletion. Expected "+
				"Delta, Geneve, moriarty, opportunity and sansa. Found %s.", n.key)
		}
	}
}

func TestSweepRemovesMarkedNodesForDeletion(t *testing.T) {
	tree := createFixedTree()
	shouldDeleteLeaves = true
	mark(tree)
	sweep(tree)
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

func TestRunningMarkAndSweepMultipleTimes(t *testing.T) {
	tree := createFixedTree()
	shouldDeleteLeaves = true
	mark(tree)
	sweep(tree)
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
	shouldDeleteLeaves = true
	mark(tree)
	sweep(tree)
	tree.shouldContain("First", t)
	tree.shouldContain("netstat", t)
	tree.shouldContain("nelly", t)
	tree.shouldContain("nefertity", t)
	tree.shouldContain("polly", t)
	tree.shouldNotContain("middle", t)
	tree.shouldNotContain("Amount", t)
}

func TestExpiredKeyShouldBeRemovedFromTree(t *testing.T) {
	tree := createFixedTree()
	tree.Insert("somewhere", "{'test': 'value_Abc'}", 500*time.Millisecond)
	time.Sleep(1 * time.Second)
	tree.purgeNodes()
	if n, _ := tree.Get("somewhere"); n != nil {
		t.Fatal("Expired node should have been removed from the tree")
	}
}

func TestNotExpiredKeyShouldNotBeRemovedFromTree(t *testing.T) {
	tree := createFixedTree()
	tree.Insert("somewhere", "{'test': 'value_Abc'}", 5*time.Second)
	time.Sleep(1 * time.Second)
	tree.purgeNodes()
	if n, _ := tree.Get("somewhere"); n == nil {
		t.Fatal("Not expired node should not have been removed from the tree")
	}
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
// 		_, _ = tree.Insert(key, "{'test': 'asdasdadasasdasdadasasdasdadasasdasdadasasdasdadasasdasdadas'}", 500*time.Second)
// 	}
// 	close(tree.jobs)
// }
