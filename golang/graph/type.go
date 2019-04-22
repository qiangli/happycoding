package graph

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
type AdjacencyList struct {
}

type IncidenceMatrix [][]int

type IncidenceList struct {
}
