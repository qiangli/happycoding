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
	sum := 21
	fmt.Printf("data: %v\nsum: %v\n", data, sum)

	b := pathsum(tree, sum)
	if !b {
		t.FailNow()
	}
}
