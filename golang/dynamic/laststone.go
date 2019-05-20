package dynamic

import (
	"fmt"
)

// https://www.geeksforgeeks.org/partition-problem-dp-18/
// https://www.geeksforgeeks.org/partition-a-set-into-two-subsets-such-that-the-difference-of-subset-sums-is-minimum/
func lastStoneWeightII(stones []int) int {
	n := len(stones)

	sum := 0
	for _, v := range stones {
		sum += v
	}

	fmt.Println("stones", stones)
	fmt.Println("sum", sum, "n", n)
	type key struct {
		i, j int
	}
	memo := make(map[key]int)

	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return x * -1
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var lsw func(int, int) int
	lsw = func(i, j int) int {
		fmt.Println("lsw", "i", i, "j", j)

		k := key{i, j}
		v, ok := memo[k]
		if ok {
			//fmt.Println("cached", "i", i, "j", j, "v", v)
			return v
		}
		if j == 0 {
			v = abs(i - (sum - i))
			fmt.Println("return", "i", i, "j", j, "v", v)
			return v
		}
		v = min(lsw(i, j-1), lsw(i+stones[j-1], j-1))
		//fmt.Println("caching", "i", i, "j", j, "v", v)
		memo[k] = v
		return v
	}

	return lsw(0, n)
}
