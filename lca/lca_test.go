package lca

import (
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	want := &node{
		val: 10,
	}
	got := newTree(10)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected %vg; got %v", want, got)
	}
}

func TestGet(t *testing.T) {
	n := newTree(10)
	insert(n, 5)
	insert(n, 11)
	insert(n, 3)

	want := &node{
		val: 11,
	}
	got := get(n, 11)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("get() failed; expected: %v, got: %v", want.val, got.val)
	}

	got = get(nil, 11)
	if got != nil {
		t.Errorf("get() failed; expected: %v, got: %v", nil, got.val)
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	// Create a binary tree to test.
	//		 10
	//		/  \
	//	 6  	13
	//  / \    \
	// 3  8 	  16
	//  \  \    /
	//  4   9  15
	root := newTree(10)
	root = insert(root, 13)
	root = insert(root, 6)
	root = insert(root, 3)
	root = insert(root, 9)
	root = insert(root, 16)
	root = insert(root, 4)
	root = insert(root, 15)
	root = insert(root, 8)

	// Case when the lca is not the root.
	a := get(root, 4)
	if a == nil {
		t.Error("requested node does not exist")
	}
	b := get(root, 9)
	if b == nil {
		t.Error("requested node does not exist")
	}
	lca := lowestCommonAncestor(root, a, b)
	if lca == nil {
		t.Error("lca does not exist")
	}
	if !reflect.DeepEqual(lca, get(root, 6)) {
		t.Errorf("expected 6; got %v", lca.val)
	}

	// Case when the root is the lca
	// and given nodes are on left and right branches
	// of the root.
	a = get(root, 16)
	if a == nil {
		t.Error("requested node does not exist")
	}
	b = get(root, 8)
	if b == nil {
		t.Error("requested node does not exist")
	}
	lca = lowestCommonAncestor(root, a, b)
	if lca == nil {
		t.Error("lca does not exist")
	}
	if !reflect.DeepEqual(lca, root) {
		t.Errorf("expected 6; got %v", lca.val)
	}

	// Case when the lca is the parent of the node itself.
	a = get(root, 15)
	if a == nil {
		t.Error("requested node does not exist")
	}
	b = get(root, 16)
	if b == nil {
		t.Error("requested node does not exist")
	}
	lca = lowestCommonAncestor(root, a, b)
	if lca == nil {
		t.Error("lca does not exist")
	}
	if !reflect.DeepEqual(lca, b) {
		t.Errorf("expected 6; got %v", lca.val)
	}

	// Case when the root is the lca.
	lca = lowestCommonAncestor(root, root, root)
	if lca == nil {
		t.Error("lca does not exist")
	}
	if !reflect.DeepEqual(lca, root) {
		t.Errorf("expected 10; got %v", lca.val)
	}

	// Case when the lowest common ancestor is itself.
	a = get(root, 4)
	if a == nil {
		t.Error("requested node does not exist")
	}
	b = get(root, 15)
	if a == nil {
		t.Error("requested node does not exist")
	}
	lca = lowestCommonAncestor(root, a, b)
	if lca == nil {
		t.Error("lca does not exist")
	}
	if !reflect.DeepEqual(lca, root) {
		t.Errorf("expected 10; got %v", lca.val)
	}

	// Case when all inuputs are nil.
	lca = lowestCommonAncestor(nil, nil, nil)
	if lca != nil {
		t.Errorf("expected: nil; got: %v", lca.val)
	}

	// Case when there is only one node.
	node := newTree(12)
	lca = lowestCommonAncestor(node, node, node)
	if lca == nil {
		t.Error("lca does not exist")
	}
	if lca != node {
		t.Errorf("expected: %v; got %v", node.val, lca.val)
	}

	// Case when there is only one node in tree
	// and input nodes are in a different tree.
	tree1 := newTree(10)
	tree2 := newTree(5)
	insert(tree2, 8)
	lca = lowestCommonAncestor(tree1, get(tree2, 5), get(tree2, 8))
	if lca != nil {
		t.Errorf("expected: nil; got: %v", lca.val)
	}
}
