package lca

import (
	"fmt"
	"testing"
)

func TestLCA(t *testing.T) {
	// Development test DFS
	g := NewGraph(10)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 6)
	g.AddEdge(3, 4)
	g.AddEdge(3, 7)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(7, 8)
	g.AddEdge(8, 9)
	g.AddEdge(9, 6)

	lca, err := g.LCA(6, 8)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Printf("%v \n", lca)

}
