package tree

func isBST(root *tree) bool {

	var check func(*tree, int, int) bool
	check = func(n *tree, lo, hi int) bool {
		if n == nil {
			return true
		}
		if n.id < lo || n.id > hi {
			return false
		}

		return check(n.left, lo, n.id-1) && check(n.right, n.id+1, hi)
	}

	if root == nil {
		return true
	}
	lo, hi := 1<<31-1, -1<<31
	var lohi func(*tree)
	lohi = func(n *tree) {
		if n == nil {
			return
		}
		if n.id < lo {
			lo = n.id
		}
		if n.id > hi {
			hi = n.id
		}
		lohi(n.left)
		lohi(n.right)
	}
	lohi(root)

	return check(root, lo, hi)
}
