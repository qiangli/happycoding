package tree

import (
	"fmt"
)

// https://www.geeksforgeeks.org/level-order-tree-traversal/
func height(n *tree) int {
	if n == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	lh := height(n.left)
	rh := height(n.right)
	h := max(lh, rh) + 1
	fmt.Println("id", n.id, "height", h, "left", lh, "right", rh)
	return h
}

func printLevelOrder(root *tree) {
	h := height(root)

	var printLevel func(*tree, int)
	printLevel = func(n *tree, level int) {
		if n == nil {
			return
		}
		if level == 1 {
			fmt.Printf(" %v ", n.id)
		} else {
			printLevel(n.left, level-1)
			printLevel(n.right, level-1)
		}
	}
	for i := 1; i <= h; i++ {
		fmt.Printf("level %v:", i)
		printLevel(root, i)
		fmt.Println()
	}
}

//
func printLevelOrder2(root *tree) {
	q := new(queue)

	q.enqueue(root)

	level := 0
	for !q.empty() {
		sz := q.size()
		level++
		fmt.Printf("level %v:", level)
		for i := 0; i < sz; i++ {
			n := q.dequeue()
			fmt.Printf(" %v ", n.id)
			if n.left != nil {
				q.enqueue(n.left)
			}
			if n.right != nil {
				q.enqueue(n.right)
			}
		}
		fmt.Println()
	}
}
