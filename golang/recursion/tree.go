package recursion

import (
	"fmt"
)

//https://en.wikipedia.org/wiki/Tree_traversal

type tree struct {
	id    int
	left  *tree
	right *tree
}

func visit(n *tree) {
	fmt.Printf("%v ", n.id)
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

func inorder(n *tree) {
	if n == nil {
		return
	}
	if n.left != nil {
		inorder(n.left)
	}

	visit(n)

	if n.right != nil {
		inorder(n.right)
	}
}

func preorder(n *tree) {
	if n == nil {
		return
	}

	visit(n)

	if n.left != nil {
		preorder(n.left)
	}
	if n.right != nil {
		preorder(n.right)
	}
}

func postorder(n *tree) {
	if n == nil {
		return
	}

	if n.left != nil {
		postorder(n.left)
	}
	if n.right != nil {
		postorder(n.right)
	}

	visit(n)

}

func outorder(n *tree) {
	if n == nil {
		return
	}

	if n.right != nil {
		outorder(n.right)
	}

	visit(n)

	if n.left != nil {
		outorder(n.left)
	}
}
