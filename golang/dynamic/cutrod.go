package dynamic

import (
	"fmt"
)

// Introduction to Algorithms - 3rd edition p363

func cutrod(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := -1 // assume min price is 0
	for i := 1; i <= n; i++ {
		q = max(q, p[i]+cutrod(p, n-i))
	}

	fmt.Printf("cut %v %v\n", n, q)
	return q
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cutrodfunctor() func([]int, int) int {
	memo := make(map[int]int)
	var f func([]int, int) int
	f = func(p []int, n int) int {
		if n == 0 {
			return 0
		}
		v, ok := memo[n]
		if ok {
			fmt.Printf("cached: %v %v\n", n, v)
			return v
		}
		//
		q := -1 //
		for i := 1; i <= n; i++ {
			q = max(q, p[i]+f(p, n-i))
		}

		fmt.Printf("caching: %v %v\n", n, q)
		memo[n] = q

		return q
	}
	return f
}
