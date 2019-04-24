package dynamic

import (
	"fmt"
)

// Introduction to Algorithms - 3rd edition p393
// longest common sequence
func lcs(x, y []byte, c [][]int, b [][]int, i, j int) int {
	max := func(a, b int) (int, int) {
		if a >= b {
			return a, 1
		}
		return b, -1
	}

	v, di := 0, 0
	if i == 0 || j == 0 {
		v = 0
	} else if x[i] == y[j] {
		v = lcs(x, y, c, b, i-1, j-1) + 1
	} else {
		v, di = max(lcs(x, y, c, b, i-1, j), lcs(x, y, c, b, i, j-1))
	}
	c[i][j] = v
	b[i][j] = di

	fmt.Printf("%v %v v: %v di: %v\n", i, j, v, di)
	return v
}

func printLcs(b [][]int, x []byte, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	if b[i][j] == 0 {
		printLcs(b, x, i-1, j-1)
		fmt.Printf("%c", x[i])
	} else if b[i][j] > 0 {
		printLcs(b, x, i-1, j)
	} else {
		printLcs(b, x, i, j-1)
	}
}
