package tree

import ()

//https://en.wikipedia.org/wiki/Tree_traversal
//iterative

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	s := NewStack()
	s.Push(root)
	for !s.IsEmpty() {
		n := s.Pop()

		visit(n)

		if n.Right != nil {
			s.Push(n.Right)
		}
		if n.Left != nil {
			s.Push(n.Left)
		}
	}
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	s := NewStack()
	n := root
	for !s.IsEmpty() || n != nil {
		if n != nil {
			s.Push(n)
			n = n.Left
		} else {
			n = s.Pop()

			visit(n)

			n = n.Right
		}
	}
}

func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	s := NewStack()
	n := root
	var last *TreeNode
	for !s.IsEmpty() || n != nil {
		if n != nil {
			s.Push(n)
			n = n.Left
		} else {
			p := s.Peek()
			if p.Right != nil && last != p.Right {
				n = p.Right
			} else {
				visit(p)
				last = s.Pop()
			}
		}
	}
}

func outOrder(root *TreeNode) {
	if root == nil {
		return
	}
	s := NewStack()
	n := root
	for !s.IsEmpty() || n != nil {
		if n != nil {
			s.Push(n)
			n = n.Right
		} else {
			n = s.Pop()
			visit(n)
			n = n.Left
		}
	}
}
