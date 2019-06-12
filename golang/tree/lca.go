package tree

import (
//"fmt"
)

//https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/submissions/
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l == nil {
		return lowestCommonAncestor(r, p, q)
	}
	if r == nil {
		return lowestCommonAncestor(l, p, q)
	}

	return root
}
