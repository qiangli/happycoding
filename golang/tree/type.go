package tree

import (
	"fmt"
)

// TreeNode is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Object *TreeNode //interface{}

type Visitor func(Object) bool

func visit(n Object) {
	fmt.Print(" -> ", n.Val)
}

// internal node for queue and stack
type node struct {
	val  Object
	next *node
}

// data is an array of arrays representing a list of tree nodes
// with format: id, left child id, right child id
func growTree(data [][3]int) *TreeNode {
	nodes := make([]TreeNode, len(data))
	find := func(Val int) *TreeNode {
		if Val == -1 {
			return nil
		}
		for i := 0; i < len(nodes); i++ {
			if nodes[i].Val == Val {
				return &nodes[i]
			}
		}
		return nil
	}
	// set value
	for i := 0; i < len(data); i++ {
		nodes[i].Val = data[i][0]
	}
	// attach child
	for i := 0; i < len(data); i++ {
		nodes[i].Left = find(data[i][1])
		nodes[i].Right = find(data[i][2])
	}
	return &nodes[0]
}
