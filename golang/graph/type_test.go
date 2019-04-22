package graph

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := new(set)

	s.add(1)
	fmt.Printf("%v\n", *s)
	s.add(2)
	fmt.Printf("%v\n", *s)

	s.add(1)
	s.add(2)
	fmt.Printf("%v\n", *s)
}

func test(g Graph, t *testing.T) {
	fmt.Printf("graph: %v\n", g)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	if !g.Adjacent(1, 2) || !g.Adjacent(1, 3) || !g.Adjacent(2, 3) {
		t.Fail()
	}
	fmt.Printf("graph: %v neighbors of 1: %v\n", g, g.Neighbors(1))

	g.RemoveEdge(2, 3)
	if g.Adjacent(2, 3) {
		t.Fail()
	}
	fmt.Printf("graph: %v\n", g)
}

func TestEdgeList(t *testing.T) {
	g := new(EdgeList)

	fmt.Println("test edge list ...")

	test(g, t)
}

func TestAdjacencyMatrix(t *testing.T) {
	g := NewAdjacencyMatrix(5)

	fmt.Println("test adjacency matrix ...")

	test(&g, t)
}
