// Package lca is the implementation for the Lowest Common Ancestor problem.
package lca

// node is a node for a binary tree.
type node struct {
	key int
	// Nodes that this node is pointing to.
	adj []*node
}

func newDigraph(key int) *node {
	return &node{
		key: key,
	}
}

func newNode(key int) *node {
	return &node{
		key: key,
	}
}

// newEdge creates an edge from node a -> b.
func newEdge(a *node, b *node) *node {
	return nil
}
