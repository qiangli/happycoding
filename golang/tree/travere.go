package tree

import (
	"fmt"
)

//https://en.wikipedia.org/wiki/Tree_traversal

type tree struct {
	id    int
	left  *tree
	right *tree
}

// simplistic stack implementation / not thread safe
type stack []*tree

func (r *stack) push(v *tree) {
	*r = append(*r, v)
	// fmt.Printf("push: %v\n", v.id)
}

func (r *stack) pop() *tree {
	size := len(*r)
	if size == 0 {
		return nil
	}
	v := (*r)[size-1]
	*r = (*r)[:size-1]
	// fmt.Printf("pop: %v\n", v.id)
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
	// fmt.Printf("empty: %v\n", len(r) == 0)
	return len(r) == 0
}

//iterative
func visit(n *tree) {
	fmt.Printf(" %v ", n.id)
}

func preorder(root *tree) {
	if root == nil {
		return
	}

	nodes := new(stack)
	nodes.push(root)
	for !nodes.empty() {
		n := nodes.pop()

		visit(n)
		if n.right != nil {
			nodes.push(n.right)
		}
		if n.left != nil {
			nodes.push(n.left)
		}
	}
}

func inorder(root *tree) {
	if root == nil {
		return
	}
	nodes := new(stack)
	n := root
	for !nodes.empty() || n != nil {
		if n != nil {
			nodes.push(n)
			n = n.left
		} else {
			n = nodes.pop()
			visit(n)
			n = n.right
		}
	}
}

func postorder(root *tree) {
	if root == nil {
		return
	}
	nodes := new(stack)
	n := root
	var last *tree
	for !nodes.empty() || n != nil {
		if n != nil {
			nodes.push(n)
			n = n.left
		} else {
			p := nodes.peek()
			if p.right != nil && last != p.right {
				n = p.right
			} else {
				visit(p)
				last = nodes.pop()
			}
		}
	}
}

func outorder(root *tree) {
	if root == nil {
		return
	}
	nodes := new(stack)
	n := root
	for !nodes.empty() || n != nil {
		if n != nil {
			nodes.push(n)
			n = n.right
		} else {
			n = nodes.pop()
			visit(n)
			n = n.left
		}
	}
}
