package tree

import (
	"fmt"
)

type tree struct {
	id    int
	left  *tree
	right *tree
}

func visit(n *tree) {
	fmt.Printf(" %v ", n.id)
}

// a is an array of arrays representing a list of tree nodes
// with format: id, left child id, right child id
func growTree(a [][3]int) *tree {
	nodes := make([]tree, len(a))
	find := func(id int) *tree {
		for i := 0; i < len(nodes); i++ {
			if nodes[i].id == id {
				return &nodes[i]
			}
		}
		return nil
	}
	// set id value
	for i := 0; i < len(a); i++ {
		nodes[i].id = a[i][0]
	}
	// attach child
	for i := 0; i < len(a); i++ {
		nodes[i].left = find(a[i][1])
		nodes[i].right = find(a[i][2])
	}
	return &nodes[0]
}

// simplistic stack implementation / not thread safe
type stack []*tree

func (r *stack) push(v *tree) {
	*r = append(*r, v)
}

func (r *stack) pop() *tree {
	size := len(*r)
	if size == 0 {
		return nil
	}
	v := (*r)[size-1]
	*r = (*r)[:size-1]
	return v
}

func (r stack) peek() *tree {
	size := len(r)
	if size == 0 {
		return nil
	}
	return r[size-1]
}

func (r stack) empty() bool {
	return len(r) == 0
}

// simplistic queue implementation / not thread safe
type queue []*tree

func (r *queue) enqueue(v *tree) {
	*r = append(*r, v)
}

func (r *queue) dequeue() *tree {
	size := len(*r)
	if size == 0 {
		return nil
	}
	v := (*r)[0]
	*r = (*r)[1:]
	return v
}

func (r queue) empty() bool {
	return len(r) == 0
}
