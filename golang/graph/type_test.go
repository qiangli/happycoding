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

func TestLinkedlist(t *testing.T) {
	ll := NewLinkedlist()

	fmt.Printf("liked list: %v\n", ll)

	ll.add(1)
	ll.add(2)
	ll.add(3)

	fmt.Printf("liked list: %v\n", ll)

	ll.remove(2)

	fmt.Printf("liked list: %v\n", ll)
	if ll.contains(2) {
		t.FailNow()
	}

	ll.remove(3)
	fmt.Printf("liked list: %v\n", ll)
	if ll.contains(3) {
		t.FailNow()
	}
}

func test(g Graph, t *testing.T) {
	fmt.Printf("test graph: %v\n", g)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	fmt.Printf("test graph: %v\n", g)

	if !g.Adjacent(1, 2) || !g.Adjacent(1, 3) || !g.Adjacent(2, 3) || !g.Adjacent(3, 4) {
		t.FailNow()
	}

	fmt.Printf("test neighbors of 1: %v\n", g.Neighbors(1))

	g.RemoveEdge(2, 3)
	fmt.Printf("test remove graph: %v\n", g)

	if g.Adjacent(2, 3) {
		t.FailNow()
	}
	fmt.Printf("test graph: %v\n", g)
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

func TestGuidoVanRossum(t *testing.T) {
	g := NewGuidoVanRossum()

	fmt.Println("test adjacency list - GuidoVanRossum ...")

	test(g, t)
}

func TestCormen(t *testing.T) {
	g := NewCormen(5)

	fmt.Println("test adjacency list - Cormen ...")

	test(g, t)
}

func TestGoodrichAndTamassia(t *testing.T) {
	g := NewGoodrichAndTamassia(5)

	fmt.Println("test adjacency list - GoodrichAndTamassia ...")

	test(g, t)
}
