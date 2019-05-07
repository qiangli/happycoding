package recursion

import (
	"fmt"
)

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	M := len(grid)
	N := len(grid[0])

	C := grid[r0][c0]

	visited := make([][]bool, M)
	for i := 0; i < M; i++ {
		visited[i] = make([]bool, N)
	}

	var visit func(int, int)

	borders := [][]int{}

	border := func(x, y int) bool {
		if x == 0 || y == 0 || x == M-1 || y == N-1 {
			return true
		}
		return grid[x-1][y] != C || grid[x+1][y] != C || grid[x][y-1] != C || grid[x][y+1] != C
	}

	check := func(x, y int) {
		if x >= 0 && x < M && y >= 0 && y < N {
			if grid[x][y] == C && !visited[x][y] {

				visit(x, y)

			}
		}
	}

	visit = func(r, c int) {

		fmt.Println("visiting", r, c, "visited", visited[r][c])
		visited[r][c] = true

		if border(r, c) {
			borders = append(borders, []int{r, c})
		}

		check(r-1, c)
		check(r+1, c)
		check(r, c-1)
		check(r, c+1)

		fmt.Println("visit", r, c, "visited", visited[r][c])
	}

	visit(r0, c0)

	for _, v := range borders {
		x := v[0]
		y := v[1]
		grid[x][y] = color
	}

	return grid
}
