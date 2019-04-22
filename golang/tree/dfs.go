package tree

//https://en.wikipedia.org/wiki/Tree_traversal
//iterative

func preorder(root *tree) {
	if root == nil {
		return
	}

	nodes := new(stack)
	nodes.push(root)
	for !nodes.empty() {
		n := nodes.pop()

		visit(n)
		if n.right != nil {
			nodes.push(n.right)
		}
		if n.left != nil {
			nodes.push(n.left)
		}
	}
}

func inorder(root *tree) {
	if root == nil {
		return
	}
	nodes := new(stack)
	n := root
	for !nodes.empty() || n != nil {
		if n != nil {
			nodes.push(n)
			n = n.left
		} else {
			n = nodes.pop()
			visit(n)
			n = n.right
		}
	}
}

func postorder(root *tree) {
	if root == nil {
		return
	}
	nodes := new(stack)
	n := root
	var last *tree
	for !nodes.empty() || n != nil {
		if n != nil {
			nodes.push(n)
			n = n.left
		} else {
			p := nodes.peek()
			if p.right != nil && last != p.right {
				n = p.right
			} else {
				visit(p)
				last = nodes.pop()
			}
		}
	}
}

func outorder(root *tree) {
	if root == nil {
		return
	}
	nodes := new(stack)
	n := root
	for !nodes.empty() || n != nil {
		if n != nil {
			nodes.push(n)
			n = n.right
		} else {
			n = nodes.pop()
			visit(n)
			n = n.left
		}
	}
}
