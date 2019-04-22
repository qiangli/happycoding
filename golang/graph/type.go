package graph

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

//
type Graph interface {
	Adjacent(x, y int) bool
	Neighbors(x int) []int
	AddEdge(x, y int)
	RemoveEdge(x, y int)
}

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

func (r *EdgeList) AddVertex(x int) {
	//
}

func (r *EdgeList) RemoveVertex(x int) {
	//
}

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

type AdjacencyMatrix [][]int

type AdjacencyList struct {
}
