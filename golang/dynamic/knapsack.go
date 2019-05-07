package dynamic

import (
	"fmt"
)

func knapsack(w int, weight []int, value []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var compute func(int, int) int
	memo := make([][]int, len(weight)+1)
	for i := 0; i < len(weight)+1; i++ {
		memo[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			memo[i][j] = -1
		}
	}
	compute = func(n int, c int) int {
		fmt.Println("n", n, "c", c)
		if n == 0 || c == 0 {
			return 0
		}

		v := memo[n][c]
		if v != -1 {
			fmt.Println("cached", n, c, v)
			return v
		}

		if weight[n-1] > c {
			return compute(n-1, c)
		}

		yes := compute(n-1, c-weight[n-1]) + value[n-1]
		no := compute(n-1, c)

		v = max(yes, no)

		fmt.Println("caching", n, c, v)
		memo[n][c] = v

		return v
	}

	v := compute(len(weight), w)
	fmt.Println(memo)
	return v
}
