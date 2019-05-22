package tree

func isBST(root *TreeNode) bool {

	var check func(*TreeNode, int, int) bool
	check = func(n *TreeNode, lo, hi int) bool {
		if n == nil {
			return true
		}
		if n.Val < lo || n.Val > hi {
			return false
		}

		return check(n.Left, lo, n.Val-1) && check(n.Right, n.Val+1, hi)
	}

	//
	lo, hi := 1<<31-1, -1<<31
	var lohi func(*TreeNode)
	lohi = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Val < lo {
			lo = n.Val
		}
		if n.Val > hi {
			hi = n.Val
		}
		lohi(n.Left)
		lohi(n.Right)
	}
	lohi(root)

	return check(root, lo, hi)
}
