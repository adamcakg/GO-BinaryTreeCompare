package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestSameTrees(t *testing.T) {
	if !Same(tree.New(1), tree.New(1)) {
		t.Fatalf("Tree same failed")
	}
}

func TestDifferentTrees(t *testing.T) {
	if Same(tree.New(1), tree.New(10)) {
		t.Fatalf("Trees are different")
	}
}
