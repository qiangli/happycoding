package array

import (
	"fmt"
)

//https://en.wikipedia.org/wiki/Maximum_subarray_problem

func maxsub(a []int) (int, int, int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	max_so_far, max_ending_here := a[0], a[0]
	old, start, end := 0, 0, 0
	for i := 1; i < len(a); i++ {
		max_ending_here = max(a[i], max_ending_here+a[i])
		max_so_far = max(max_so_far, max_ending_here)
		fmt.Printf("ending: %v %v so far: %v\n", i, max_ending_here, max_so_far)
		if max_ending_here < 0 {
			start = i + 1
		}
		if max_ending_here == max_so_far {
			old = start
			end = i
		}
	}
	return old, end, max_so_far
}
