// Package lca is the implementation for the Lowest Common Ancestor problem.
package lca

import (
	"fmt"
)

// Digraph is a directed graph represented by an adjacency matrix, e.g.:
// A digraph with 5 nodes, with node u as rows and node v as columns, and
// [u][v] = T is an edge from u -> v.
//
//		| 0 | 1 | 2 | 3 | 4
//	0 | t | t | t | f | f
//  1 | f | t | t | f | f
//	2 | f | f | t | t | t
//	3 | f | f | f | t | t
//	4 | f | f | f | f | t
type Digraph struct {
	nodeCount int
	edges     [][]bool
}

// AddEdge adds an edge to a digraph.
func (g *Digraph) AddEdge(u, v int) error {
	err := g.isInGraph(u, v)
	if err != nil {
		return err
	}
	g.edges[u][v] = true
	return nil
}

// RemoveEdge removes an edge from a digraph.
func (g *Digraph) RemoveEdge(u, v int) error {
	err := g.isInGraph(u, v)
	if err != nil {
		return err
	}
	g.edges[u][v] = false
	return nil
}

func (g *Digraph) isInGraph(u, v int) error {
	if u > g.nodeCount || u < 0 {
		return fmt.Errorf("node %v is not in the graph", u)
	}
	if v > g.nodeCount || v < 0 {
		return fmt.Errorf("node %v is not in the graph", v)
	}
	return nil
}

// LCA finds the lowest common ancestor of nodes n and m.
func LCA(n, m int) (int, error) {

	return 0, nil
}
