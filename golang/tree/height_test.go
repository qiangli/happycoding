package tree

import (
	"fmt"
	"testing"
)

func TestHeight(t *testing.T) {
	data := [][3]int{
		{4, 2, 5},
		{2, 1, 3},
		{1, -1, -1},
		{3, -1, -1},
		{5, -1, 6},
		{6, -1, 7},
		{7, -1, -1},
	}
	tree := growTree(data)

	fmt.Printf("data: %v\n", data)

	h := height(tree)
	fmt.Println("height", h)
	if h != 4 {
		t.FailNow()
	}

	printLevelOrder(tree)
}
