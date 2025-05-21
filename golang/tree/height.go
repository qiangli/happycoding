package tree

import (
	"fmt"
)

// https://www.geeksforgeeks.org/level-order-tree-traversal/
func height(n *TreeNode) int {
	if n == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	lh := height(n.Left)
	rh := height(n.Right)
	h := max(lh, rh) + 1
	fmt.Println("id", n.Val, "height", h, "left", lh, "right", rh)
	return h
}

func printLevelOrder(root *TreeNode) {
	h := height(root)

	var printLevel func(*TreeNode, int)
	printLevel = func(n *TreeNode, level int) {
		if n == nil {
			return
		}
		if level == 1 {
			fmt.Printf(" %v ", n.Val)
		} else {
			printLevel(n.Left, level-1)
			printLevel(n.Right, level-1)
		}
	}
	for i := 1; i <= h; i++ {
		fmt.Printf("level %v:", i)
		printLevel(root, i)
		fmt.Println()
	}
}

func printLevelOrder2(root *TreeNode) {
	q := NewQueue()

	q.Enqueue(root)

	level := 0
	for !q.IsEmpty() {
		sz := q.Size()
		level++
		fmt.Printf("level %v:", level)
		for i := 0; i < sz; i++ {
			n := q.Dequeue()
			fmt.Printf(" %v ", n.Val)
			if n.Left != nil {
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				q.Enqueue(n.Right)
			}
		}
		fmt.Println()
	}
}
