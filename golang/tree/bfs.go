package tree

import ()

//https://en.wikipedia.org/wiki/Tree_traversal
// breadth first search

func levelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	q := NewQueue()
	q.Enqueue(root)
	for !q.IsEmpty() {
		n := q.Dequeue()

		visit(n)

		if n.Left != nil {
			q.Enqueue(n.Left)
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
		}
	}
}
