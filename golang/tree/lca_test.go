package tree

import (
	"fmt"
	"testing"
)

func TestLca(t *testing.T) {
	//[3,5,1,6,2,0,8,null,null,7,4]
	data := [][3]int{
		{3, 5, 1},
		{5, 6, 2},
		{6, -1, -1},
		{2, 7, 4},
		{7, -1, -1},
		{4, -1, -1},
		{1, 0, 8},
		{0, -1, -1},
		{8, -1, -1},
	}
	tree := growTree(data)
	//5, 4
	var find func(*TreeNode, int) *TreeNode
	find = func(n *TreeNode, v int) *TreeNode {
		if n == nil {
			return nil
		}
		if n.Val == v {
			return n
		}
		l := find(n.Left, v)
		if l != nil {
			return l
		}
		r := find(n.Right, v)
		return r
	}
	p := find(tree, 5)
	q := find(tree, 4)

	n := lowestCommonAncestor(tree, p, q)
	fmt.Println("lca", n)
}
