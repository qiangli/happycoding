//https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/995/
package tree

import (
	"fmt"
	"strconv"
	"strings"
)

func serialize(root *TreeNode) string {
	if root == nil {
		return "-"
	}
	//
	l := serialize(root.Left)
	r := serialize(root.Right)

	return fmt.Sprintf("%v,%v,%v", root.Val, l, r)
}

func deserialize(data string) *TreeNode {
	var build func([]string, int) (*TreeNode, int)
	build = func(a []string, idx int) (*TreeNode, int) {
		idx++
		s := a[idx]

		if s == "-" {
			return nil, idx
		}

		v, _ := strconv.Atoi(s)
		n := &TreeNode{Val: v}
		n.Left, idx = build(a, idx)
		n.Right, idx = build(a, idx)

		return n, idx
	}
	a := strings.Split(data, ",")
	n, _ := build(a, -1)
	return n
}

func buildTree(data string) *TreeNode {
	sa := strings.Split(data, ",")
	n := len(sa)
	nodes := make([]*TreeNode, n)
	//
	for i, s := range sa {
		if s == "null" {
			continue
		}
		v, _ := strconv.Atoi(s)
		nodes[i] = &TreeNode{Val: v}
	}
	//
	for i, j := 0, 0; i < n && 2*j < n; i++ {
		if nodes[i] != nil {
			if 2*j+1 < n {
				nodes[i].Left = nodes[2*j+1]
			}
			if 2*j+2 < n {
				nodes[i].Right = nodes[2*j+2]
			}
			j++
		}
	}
	return nodes[0]
}
