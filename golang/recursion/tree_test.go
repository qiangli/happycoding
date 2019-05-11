package recursion

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

func TestBstTiGst(t *testing.T) {
	data := [][3]int{
		{4, 1, 6},
		{1, 0, 2},
		{0, -1, -1},
		{2, -1, 3},
		{3, -1, -1},
		{6, 5, 7},
		{5, -1, -1},
		{7, -1, 8},
		{8, -1, -1},
	}
	tree := growTree(data)

	fmt.Printf("data: %v\n", data)

	fmt.Println("\nbst tree")
	bstToGst(tree)

	inorder(tree)

	fmt.Println()
}
