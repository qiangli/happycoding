package tree

import (
	"fmt"
)

// https://www.geeksforgeeks.org/root-to-leaf-path-sum-equal-to-a-given-number/
func pathsum(root *TreeNode, sum int) bool {
	var compute func(*TreeNode, int, int) bool
	compute = func(n *TreeNode, sum, total int) bool {
		leaf := n.Left == nil && n.Right == nil
		total += n.Val
		fmt.Println("leaf", leaf, "sum", sum, "total", total)

		if leaf {
			return total == sum
		}

		return compute(n.Left, sum, total) || compute(n.Right, sum, total)
	}

	return compute(root, sum, 0)
}
