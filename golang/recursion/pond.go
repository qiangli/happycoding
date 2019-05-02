package recursion

//compute pond size

func matrix(row, col int) [][]int {
	var m [][]int
	m = make([][]int, row)
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, col)
	}
	return m
}

func computePondSize(land [][]int) []int {
	n := len(land)
	visited := matrix(n, n)
	sizes := []int{}

	var compute func(int, int) int
	compute = func(r, c int) int {
		s := 1
		visited[r][c] = 1
		for _, i := range []int{r - 1, r, r + 1} {
			for _, j := range []int{c - 1, c, c + 1} {
				if i >= 0 && i < n &&
					j >= 0 && j < n &&
					land[i][j] == 0 &&
					visited[i][j] == 0 {
					s += compute(i, j)
				}
			}
		}
		return s
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 0 && visited[i][j] == 0 {
				sz := compute(i, j)
				sizes = append(sizes, sz)
			}
		}
	}
	return sizes
}
