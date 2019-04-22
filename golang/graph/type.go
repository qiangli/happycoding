package graph

import (
	"fmt"
	"strconv"
)

// https://en.wikipedia.org/wiki/Graph_(abstract_data_type)
// AdjacencyList
// AdjacencyMatrix
// IncidenceMatrix
//

// simplistic set implementation
type set []int

func (r *set) add(x int) bool {
	for _, v := range *r {
		if x == v {
			return false
		}
	}
	*r = append(*r, x)
	return true
}

// simplistic linked list implementation
type node struct {
	val  int
	next *node
}

type linkedlist struct {
	head *node
}

func (r linkedlist) String() string {
	var s string
	for i := r.head.next; ; i = i.next {
		if i == nil {
			return s + " ->nil"
		}
		s = s + " -->" + strconv.Itoa(i.val)
	}
}

func NewLinkedlist() *linkedlist {
	head := &node{
		val: -1,
	}
	return &linkedlist{head}
}

func (r *linkedlist) add(v int) {
	c := r.head
	for c.next != nil {
		if c.val == v {
			return
		}
		c = c.next
	}
	c.next = &node{
		val: v,
	}
}

func (r *linkedlist) remove(v int) {
	h := r.head
	for p, c := h, h.next; ; {
		if c.val == v {
			p.next = c.next
			return
		}
		if c.next == nil {
			return
		}
		p = c
		c = c.next
	}
}

func (r *linkedlist) contains(v int) bool {
	h := r.head
	if h.next == nil {
		return false
	}
	for i := h.next; i.next != nil; i = i.next {
		if i.val == v {
			return true
		}
	}
	return false
}

//
type Graph interface {
	Adjacent(x, y int) bool
	Neighbors(x int) []int
	AddEdge(x, y int)
	RemoveEdge(x, y int)
	AddVertex(x int)
	RemoveVertex(x int)
}

//
type EdgeList [][2]int

func (r EdgeList) Adjacent(x, y int) bool {
	for _, v := range r {
		if x == v[0] && y == v[1] || x == v[1] && y == v[0] {
			return true
		}
	}
	return false
}

func (r EdgeList) Neighbors(x int) []int {
	s := new(set)

	for _, v := range r {
		if x == v[0] {
			s.add(v[1])
		} else if x == v[1] {
			s.add(v[0])
		}
	}
	return []int(*s)
}

func (r *EdgeList) AddVertex(x int) {}

func (r *EdgeList) RemoveVertex(x int) {}

func (r *EdgeList) AddEdge(x, y int) {
	if r.Adjacent(x, y) {
		return
	}
	v := [2]int{x, y}
	*r = append(*r, v)
}

func (r *EdgeList) RemoveEdge(x, y int) {
	for i, v := range *r {
		if x == v[0] && y == v[1] || x == v[1] && y == v[0] {
			*r = append((*r)[0:i], (*r)[i+1:]...)
			return
		}
	}
}

// AdjacencyMatrix is a two-dimensional matrix, in which
// the rows represent source vertices and columns represent destination vertices.
// Data on edges and vertices must be stored externally.
type AdjacencyMatrix [][]int

func NewAdjacencyMatrix(n int) AdjacencyMatrix {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func (r AdjacencyMatrix) Adjacent(x, y int) bool {
	return r[x][y] == 1 || r[y][x] == 1
}

func (r AdjacencyMatrix) Neighbors(x int) []int {
	s := new(set)
	for i, v := range r[x] {
		if v == 1 {
			s.add(i)
		}
	}
	return []int(*s)
}

func (r *AdjacencyMatrix) AddEdge(x, y int) {
	(*r)[x][y] = 1
	(*r)[y][x] = 1
}

func (r *AdjacencyMatrix) RemoveEdge(x, y int) {
	(*r)[x][y] = 0
	(*r)[y][x] = 0
}

func (r *AdjacencyMatrix) AddVertex(x int)    {}
func (r *AdjacencyMatrix) RemoveVertex(x int) {}

// https://en.wikipedia.org/wiki/Adjacency_list
// Vertices are stored as records or objects, and every vertex stores a list of adjacent vertices

// GuidoVanRossum suggested a hash table to associate each vertex in a graph
// with an array of adjacent vertices
type GuidoVanRossum struct {
	hashtable map[int][]int
}

func NewGuidoVanRossum() *GuidoVanRossum {
	return &GuidoVanRossum{
		hashtable: map[int][]int{},
	}
}

func (r *GuidoVanRossum) Adjacent(x, y int) bool {
	for _, v := range r.hashtable[x] {
		if v == y {
			return true
		}
	}
	return false
}

func (r *GuidoVanRossum) Neighbors(x int) []int {
	return r.hashtable[x]
}

func (r *GuidoVanRossum) AddEdge(x, y int) {
	a, found := r.hashtable[x]
	if !found {
		a = []int{y}
	} else {
		a = append(a, y)
	}
	r.hashtable[x] = a
}

func (r *GuidoVanRossum) RemoveEdge(x, y int) {
	a, found := r.hashtable[x]
	if found {
		for i := 0; i < len(a); i++ {
			if a[i] == y {
				r.hashtable[x] = append(a[0:i], a[i+1:]...)
				return
			}
		}
	}
}

func (r *GuidoVanRossum) AddVertex(x int)    {}
func (r *GuidoVanRossum) RemoveVertex(x int) {}

// Cormen et al. suggest an implementation in which the vertices are represented by index numbers.
// Their representation uses an array indexed by vertex number, in which the array cell for each vertex
// points to a singly linked list of the neighboring vertices of that vertex.

type Cormen struct {
	array []*linkedlist
}

func (r Cormen) String() string {
	var s string
	s = "["
	for i, v := range r.array {
		s += fmt.Sprintf(" %v: %v,", i, v)
	}
	s += "]"
	return s
}

func NewCormen(n int) *Cormen {
	return &Cormen{
		array: make([]*linkedlist, n),
	}
}

func (r *Cormen) Adjacent(x, y int) bool {
	l := r.array[x]
	if l.head.next == nil {
		return false
	}
	for i := l.head.next; ; {
		if i.val == y {
			return true
		}
		if i.next == nil {
			break
		}
		i = i.next
	}
	return false
}

func (r *Cormen) Neighbors(x int) []int {
	var a []int
	h := r.array[x].head
	if h.next == nil {
		return a
	}
	for i := h.next; ; {
		a = append(a, i.val)
		if i.next == nil {
			break
		}
		i = i.next
	}
	return a
}

func (r *Cormen) AddEdge(x, y int) {
	l := r.array[x]
	if l == nil {
		l = NewLinkedlist()
		r.array[x] = l
	}
	l.add(y)
	//
	l = r.array[y]
	if l == nil {
		l = NewLinkedlist()
		r.array[y] = l
	}
	l.add(x)
}

func (r *Cormen) RemoveEdge(x, y int) {
	l := r.array[x]
	l.remove(y)
	//
	l = r.array[y]
	l.remove(x)
}

func (r *Cormen) AddVertex(x int) {

}

func (r *Cormen) RemoveVertex(x int) {

}

//
type object interface{}

type vertex struct {
	val  int
	edge []*edge
}

func (r vertex) String() string {
	return fmt.Sprintf("%v", r.val)
}

func (r *vertex) addEdge(e *edge) {
	x := e.vertex[0].val
	y := e.vertex[1].val

	for _, v := range r.edge {
		if v != nil && sameEdge(v, x, y) {
			return
		}
	}
	r.edge = append(r.edge, e)
}

func (r *vertex) removeEdge(e *edge) {
	x := e.vertex[0].val
	y := e.vertex[1].val

	for i, v := range r.edge {
		if v != nil && sameEdge(v, x, y) {
			r.edge = append(r.edge[0:i], r.edge[i+1:]...)
		}
	}
}

func sameEdge(e *edge, x, y int) bool {
	return e.vertex[0].val == x && e.vertex[1].val == y || e.vertex[1].val == x && e.vertex[0].val == y
}

type edge struct {
	vertex [2]*vertex
}

func (r edge) String() string {
	return fmt.Sprintf("%v-%v", r.vertex[0], r.vertex[1])
}

// GoodrichAndTamassia incidence list has special classes of vertex objects and edge objects.
// Each vertex object has an instance variable pointing to a collection object
// that lists the neighboring edge objects. In turn, each edge object points to the two vertex objects
// at its endpoints
// https://github.com/fletchto99/CSI2110/blob/master/Labs/resources/net/datastructures/AdjacencyMapGraph.java
type GoodrichAndTamassia struct {
	vertices []*vertex
	edges    []*edge
}

func (r GoodrichAndTamassia) String() string {
	var s string
	s = "{ vertices: ["
	for _, v := range r.vertices {
		s += fmt.Sprintf("%v, ", v)
	}
	s += "], edges: ["
	for _, v := range r.edges {
		s += fmt.Sprintf("%v, ", v)
	}
	s += "]}"
	return s
}

func NewGoodrichAndTamassia(n int) *GoodrichAndTamassia {
	return &GoodrichAndTamassia{
		vertices: make([]*vertex, n),
	}
}

func (r *GoodrichAndTamassia) Adjacent(x, y int) bool {
	for _, e := range r.edges {
		if sameEdge(e, x, y) {
			return true
		}
	}
	return false
}

func (r *GoodrichAndTamassia) Neighbors(x int) []int {
	var a []int
	vx := r.vertices[x]
	if vx == nil {
		return a
	}
	for _, e := range vx.edge {
		if e != nil {
			for _, v := range e.vertex {
				if v != nil && v.val != x {
					a = append(a, v.val)
				}
			}
		}
	}
	return a
}

func (r *GoodrichAndTamassia) AddEdge(x, y int) {
	vx := r.vertex(x)
	vy := r.vertex(y)
	xy := r.edge(x, y)

	vx.addEdge(xy)
	vy.addEdge(xy)
}

func (r *GoodrichAndTamassia) RemoveEdge(x, y int) {
	vx := r.vertex(x)
	vy := r.vertex(y)
	xy := r.edge(x, y)

	vx.removeEdge(xy)
	vy.removeEdge(xy)

	for i, e := range r.edges {
		if sameEdge(e, x, y) {
			r.edges = append(r.edges[0:i], r.edges[i+1:]...)
		}
	}
}

// find and return the vertex; create if it does not exist
func (r *GoodrichAndTamassia) vertex(x int) *vertex {
	v := r.vertices[x]
	if v == nil {
		v = &vertex{
			val: x,
		}
		r.vertices[x] = v
	}
	return v
}

// find and return the edge; create if it does not exist
func (r *GoodrichAndTamassia) edge(x, y int) *edge {
	var xy *edge
	for _, e := range r.edges {
		if sameEdge(e, x, y) {
			xy = e
			break
		}
	}
	if xy == nil {
		vx := r.vertex(x)
		vy := r.vertex(y)
		xy = &edge{
			vertex: [2]*vertex{vx, vy},
		}
		r.edges = append(r.edges, xy)
	}
	return xy
}

func (r *GoodrichAndTamassia) AddVertex(x int) {}

func (r *GoodrichAndTamassia) RemoveVertex(x int) {}

//
type IncidenceMatrix [][]int

//
type IncidenceList struct {
}
