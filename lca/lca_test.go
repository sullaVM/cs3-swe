package lca

import (
	"fmt"
	"testing"
)

func TestAddEdge(t *testing.T) {
	g := NewGraph(3)

	err := g.AddEdge(0, 5)
	if err == nil {
		t.Error("removing non-existent edge failed")
	}

	err = g.AddEdge(0, 1)
	if err != nil {
		t.Error("removing non-existent edge failed")
	}

}

func TestRemoveEdge(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)

	err := g.RemoveEdge(5, 4)
	if err == nil {
		t.Error("removing non-existent edge failed")
	}

	err = g.RemoveEdge(1, 4)
	if err == nil {
		t.Error("removing non-existent edge failed")
	}

	err = g.RemoveEdge(0, 1)
	if err != nil {
		t.Error("removing non-existent edge failed")
	}
}

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

	// Case 1:
	// The LCA is directly a step away from both.
	got, err := g.LCA(4, 7)
	want := 3
	if err != nil {
		fmt.Print(err)
	}
	if want != got {
		t.Errorf("LCA failed; got: %v; want: %v", got, want)
	}

	// Case 2:
	// The LCA of nodes u and v is u itself and there are multiple
	// other common ancestors of the 2.
	got, err = g.LCA(6, 8)
	want = 8
	if err != nil {
		fmt.Print(err)
	}
	if want != got {
		t.Errorf("LCA failed; got: %v; want: %v", got, want)
	}

	// Case 3:
	// The LCA of nodes u and v is itself if u == v.
	got, err = g.LCA(0, 0)
	want = 0
	if err != nil {
		fmt.Print(err)
	}
	if want != got {
		t.Errorf("LCA failed; got: %v; want: %v", got, want)
	}

	// Case 4:
	// When there is no LCA for nodes u and v, it should return -1 and an error.
	g.RemoveEdge(7, 8)
	got, err = g.LCA(8, 5)
	want = -1
	if err != nil {
		fmt.Print(err)
	}
	if want != got {
		t.Errorf("LCA failed; got: %v; want: %v", got, want)
	}

	// Case 5:
	// The LCA when the distance from the LCA to node u is less than
	// the distance from the LCA to node v.
	got, err = g.LCA(5, 7)
	want = 3
	if err != nil {
		fmt.Print(err)
	}
	if want != got {
		t.Errorf("LCA failed; got: %v; want: %v", got, want)
	}

	// Case 6:
	// The LCA when there is 2 nodes that are both LCA of nodes u and v.
	got, err = g.LCA(6, 1)
	want1 := 0
	want2 := 1
	if err != nil {
		fmt.Print(err)
	}
	if want1 != got && want2 != got {
		t.Errorf("LCA failed; got: %v; want: %v", got, want)
	}
}
