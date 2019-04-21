package recursion

import (
	"fmt"
	"testing"
)

func TestInorder(t *testing.T) {
	data := [][3]int{
		{4, 2, 5},
		{2, 1, 3},
		{1, -1, -1},
		{3, -1, -1},
		{5, -1, 6},
		{6, -1, -1},
	}
	tree := growTree(data)

	fmt.Printf("data: %v\n", data)

	fmt.Println("\npre order")
	preorder(tree)

	fmt.Println("\nin order")
	inorder(tree)

	fmt.Println("\npost order")
	postorder(tree)

	fmt.Println("\nout order")
	outorder(tree)

	fmt.Println()
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
