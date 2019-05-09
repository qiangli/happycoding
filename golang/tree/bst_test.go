package tree

import (
	"testing"
)

func TestIsBST(t *testing.T) {
	data := [][3]int{
		{4, 2, 5},
		{2, 1, 3},
		{1, -1, -1},
		{3, -1, -1},
		{5, -1, 6},
		{6, -1, -1},
	}
	tree := growTree(data)

	b := isBST(tree)
	if !b {
		t.FailNow()
	}

	data = [][3]int{
		{2, 4, 5},
		{4, 1, 3},
		{1, -1, -1},
		{3, -1, -1},
		{5, -1, 6},
		{6, -1, -1},
	}

	tree = growTree(data)

	b = isBST(tree)
	if b {
		t.FailNow()
	}
}
