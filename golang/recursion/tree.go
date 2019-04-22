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
