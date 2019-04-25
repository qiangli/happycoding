package dynamic

import (
	"fmt"
)

//
// greedy algorithms p419

func activitySelect(s, f []int, k, n int) []int {
	i := k + 1
	for ; i <= n && s[i] < f[k]; i++ {
		//
	}
	if i <= n {
		fmt.Printf("activity sel: %v %v\n", i, k)
		a := append([]int{i}, activitySelect(s, f, i, n)...)
		fmt.Printf("activity sel: %v %v %v\n", i, k, a)
		return a
	}
	return []int{}
}
