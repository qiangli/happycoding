package tree

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
	data := [][3]int{
		{4, 2, 5},
		{2, 1, 3},
		{1, -1, -1},
		{3, -1, -1},
		{5, -1, 6},
		{6, -1, -1},
	}
	tree := growTree(data)

	fmt.Printf("data: %v\n", data)

	fmt.Println("\npre order")
	preorder(tree)

	fmt.Println("\nin order")
	inorder(tree)

	fmt.Println("\npost order")
	postorder(tree)

	fmt.Println("\nout order")
	outorder(tree)

	fmt.Println()
}
