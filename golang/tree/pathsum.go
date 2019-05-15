package tree

import (
	"fmt"
)

// https://www.geeksforgeeks.org/root-to-leaf-path-sum-equal-to-a-given-number/
func pathsum(root *tree, sum int) bool {
	var compute func(*tree, int, int) bool
	compute = func(n *tree, sum, total int) bool {
		leaf := n.left == nil && n.right == nil
		total += n.id
		fmt.Println("leaf", leaf, "sum", sum, "total", total)

		if leaf {
			return total == sum
		}

		return compute(n.left, sum, total) || compute(n.right, sum, total)
	}

	return compute(root, sum, 0)
}
