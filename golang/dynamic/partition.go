package dynamic

import (
	"fmt"
)

// https://www.geeksforgeeks.org/partition-problem-dp-18/
func partition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	type key struct {
		n int
		s int
	}
	memo := make(map[key]bool)
	var dp func(int, int) bool
	dp = func(n int, s int) bool {
		if s == 0 {
			return true
		}
		if n == 0 {
			return false
		}

		v, ok := memo[key{n, s}]
		if ok {
			fmt.Println("cached", v)
			return v
		}

		if nums[n-1] > sum {
			v = dp(n-1, s)
		} else {
			v = dp(n-1, s) || dp(n-1, s-nums[n-1])
		}
		fmt.Println("caching", v)
		memo[key{n, s}] = v
		return v
	}
	res := dp(len(nums), sum/2)
	return res
}

func partition2(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	n := len(nums)
	p := make([][]bool, sum/2+1)
	for i := 0; i < len(p); i++ {
		p[i] = make([]bool, n+1)
	}

	for j := 0; j <= n; j++ {
		p[0][j] = true
	}
	for i := 0; i <= sum/2; i++ {
		p[i][0] = false
	}

	for i := 1; i <= sum/2; i++ {
		for j := 1; j <= n; j++ {
			p[i][j] = p[i][j-1]
			if i >= nums[j-1] {
				p[i][j] = p[i][j] || p[i-nums[j-1]][j-1]
			}
		}
	}
	return p[sum/2][n]
}
