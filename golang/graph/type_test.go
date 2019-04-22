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

func TestEdgeList(t *testing.T) {
	g := new(EdgeList)
	fmt.Printf("edge list: %v\n", g)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	if !g.Adjacent(1, 2) || !g.Adjacent(1, 3) || !g.Adjacent(2, 3) {
		t.Fail()
	}
	fmt.Printf("edge list: %v neighbors of 1: %v\n", g, g.Neighbors(1))

	g.RemoveEdge(2, 3)
	if g.Adjacent(2, 3) {
		t.Fail()
	}
	fmt.Printf("edge list: %v\n", g)
}
