package tree

import (
	"fmt"
	"testing"
)

func TestPathsum(t *testing.T) {
	data := [][3]int{
		{10, 8, 2},
		{8, 3, 5},
		{3, -1, -1},
		{5, -1, -1},
		{2, 2, -1},
		{2, -1, -1},
	}
	tree := growTree(data)

	fmt.Printf("data: %v\n", data)

	b := pathsum(tree, 21)
	if !b {
		t.FailNow()
	}
}
