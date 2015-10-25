package splay

import (
	//"fmt"
	"testing"
)

func createFixedTree() *Tree {
	fakeTree := New(1)
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
			t.Fatalf("Wrong node marked for deletion. Expected one between"+
				"Delta, Geneve, moriarty, opportunity and sansa. Found %s.", n.key)
		}
	}
}

func TestSweepRemovesMarkedNodesForDeletion(t *testing.T) {
	tree := createFixedTree()
	mark(tree)
	sweep(tree)
	// Check we have the good nodes remaining
}
