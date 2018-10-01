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

func TestLowestCommonAncestor(t *testing.T) {
	// Create a binary tree to test.
	//		 10
	//		/ \
	//	 6  13
	//  / \   \
	// 3  9   16
	//  \    /
	//  4   15
	root := newTree(10)
	root = insert(root, 13)
	root = insert(root, 6)
	root = insert(root, 3)
	root = insert(root, 9)
	root = insert(root, 16)
	root = insert(root, 4)
	root = insert(root, 15)

	// Normal case.
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

	// Case when the root is the lca.
	lca = lowestCommonAncestor(root, root, root)
	if lca == nil {
		t.Error("lca does not exist")
	}
	if !reflect.DeepEqual(lca, root) {
		t.Errorf("expected 10; got %v", lca)
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
		t.Errorf("expected 10; got %v", lca)
	}

}
