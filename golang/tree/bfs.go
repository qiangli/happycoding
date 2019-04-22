package tree

//https://en.wikipedia.org/wiki/Tree_traversal
// breadth first search

func levelorder(root *tree) {
	if root == nil {
		return
	}

	nodes := new(queue)
	nodes.enqueue(root)
	for !nodes.empty() {
		n := nodes.dequeue()
		visit(n)
		if n.left != nil {
			nodes.enqueue(n.left)
		}
		if n.right != nil {
			nodes.enqueue(n.right)
		}
	}
}
