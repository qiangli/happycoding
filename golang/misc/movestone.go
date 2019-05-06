package misc

import (
	"fmt"
	"sort"
)

//https://blog.csdn.net/zjucor/article/details/89844621
func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)

	fmt.Println(stones)
	fmt.Println(len(stones))

	n := len(stones)

	lo := n
	for i, j := 0, 0; j < n; j++ {
		for stones[j]-stones[i] >= n {
			i++
		}
		if j-i+1 == n-1 && stones[j]-stones[i] == n-2 {
			lo = Min(lo, 2)
		} else {
			lo = Min(lo, n-(j-i+1))
		}
		fmt.Println("i", i, "j", j, " lo:", lo)
	}

	hi := Max(stones[n-1]-stones[1], stones[n-2]-stones[0]) - (n - 2)
	return []int{lo, hi}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
