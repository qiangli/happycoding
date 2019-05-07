package recursion

import (
	"fmt"
	"testing"
)

func TestColorBorder(t *testing.T) {
	grid := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	colorBorder(grid, 1, 1, 2)
	fmt.Println(grid)

	grid = [][]int{{2, 2, 2, 3, 3}, {2, 3, 3, 1, 2}, {2, 1, 3, 2, 1}}
	colorBorder(grid, 1, 4, 1)
	fmt.Println(grid)

	grid = [][]int{{1, 2, 1, 2, 1, 2}, {2, 2, 2, 2, 1, 2}, {1, 2, 2, 2, 1, 2}}
	colorBorder(grid, 1, 3, 1)
	fmt.Println(grid)
}
