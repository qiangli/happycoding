package dynamic

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Factorial

func factorial(n int) int {
	fmt.Printf("calling: %v\n", n)

	if n == 0 {
		return 1
	}

	r := factorial(n-1) * n

	fmt.Printf("factorial %v %v\n", n, r)
	return r
}

func facfunctor() func(int) int {
	memo := make(map[int]int)
	var f func(int) int

	f = func(n int) int {
		if n == 0 {
			return 1
		}
		v, ok := memo[n]
		if ok {
			fmt.Printf("cached: %v %v\n", n, v)
			return v
		}
		v = f(n-1) * n
		fmt.Printf("caching: %v %v\n", n, v)
		memo[n] = v

		return v
	}
	return f
}
