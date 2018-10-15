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
	visited   []bool
	min       int
}

// NewGraph returns a new digraph isntance.
func NewGraph(n int) *Digraph {
	ed := make([][]bool, n)
	for i := range ed {
		ed[i] = make([]bool, n)
	}
	return &Digraph{
		nodeCount: n,
		edges:     ed,
		visited:   make([]bool, n),
		min:       n,
	}
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

// DFS finds the given node 'find' and counts the steps it took to get there.
func (g *Digraph) DFS(i, find, steps int) int {
	if find == i {
		fmt.Printf("%v -> %v %v\n", i, find, steps)
		if steps < g.min {
			g.min = steps
		}
		return steps
	}
	steps++
	// Initialise tmpSteps to -1, to mean 'find' is not found.
	tmpSteps := -1
	for j := 0; j < g.nodeCount; j++ {
		if !g.visited[j] && g.edges[i][j] {
			tmpSteps = g.DFS(j, find, steps)
			fmt.Printf("%v -> %v %v\n", i, j, tmpSteps)
		}
	}
	if tmpSteps < g.min && tmpSteps != -1 {
		g.min = tmpSteps
	}
	return tmpSteps
}

func min(arr []int) int {
	min := arr[0]
	for _, i := range arr {
		if min > i {
			min = i
		}
	}
	return min
}

// LCA finds the lowest common ancestor of nodes n and m.
func (g *Digraph) LCA(n, m int) (int, error) {
	// If n == m, the LCA is automatically itself so it's not necessary
	// to evaluate.
	if n == m {
		return 0, nil
	}

	nodeDistanceTable := map[int]int{}
	for i := 0; i < g.nodeCount; i++ {
		g.resetMin()
		g.DFS(i, n, 0)
		// Set distance from i -> n.
		// If there is no path from i -> n, don't add this to the distance table.
		if g.min == g.nodeCount {
			continue
		}
		nodeDistanceTable[i] = g.min
		fmt.Printf("n min: %v\n", g.min)

		g.resetMin()
		g.DFS(i, m, 0)
		// If the distance of i -> m is greater than
		// i -> n, make this the 'real' distance from i -> n, m.
		fmt.Printf("m min: %v\n\n", g.min)
		if nodeDistanceTable[i] < g.min && g.min != g.nodeCount {
			nodeDistanceTable[i] = g.min
		}
		// If there is no path from i -> m, then make the min path -1.
		if g.min == g.nodeCount {
			delete(nodeDistanceTable, i)
		}
	}

	// Set the min to the maximum amount of directed edges graph g can have.
	min := g.nodeCount * 2
	minNode := -1
	for k, v := range nodeDistanceTable {
		// If v < 0, then v is either 0 where the node is pointing
		// to itself or -1 where the there means the given node k is
		// not a common ancestor at all.
		if min > v && v > 0 {
			min = v
			minNode = k
		}
		fmt.Printf("%v -> %v\n", k, v)
	}
	fmt.Print("\n---- \n")

	if minNode == -1 {
		return -1, fmt.Errorf("no common ancestor for %v and %v", n, m)
	}

	return minNode, nil
}
func (g *Digraph) resetMin() {
	g.min = g.nodeCount
}
