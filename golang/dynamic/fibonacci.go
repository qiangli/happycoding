package dynamic

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Fibonacci_number

func fibonacci(n int) int {
	fmt.Printf("calling %v\n", n)

	if n == 0 || n == 1 {
		return n
	}

	fn := fibonacci(n-1) + fibonacci(n-2)

	fmt.Printf("fibonacci: %v %v\n", n, fn)
	return fn
}

func fibfunctor() func(int) int {
	memo := make(map[int]int)
	var hit int
	var fib func(int) int
	fib = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		v, ok := memo[n]
		if ok {
			hit++
			fmt.Printf("cached: %v %v hit: %v\n", n, v, hit)
			return v
		}
		v = fib(n-1) + fib(n-2)

		fmt.Printf("caching: %v %v\n", n, v)
		memo[n] = v
		return v
	}

	return fib
}

// in O(n) time and O(1) space
func fibonacciN1(n int) int {
	// var memo []int = make([]int, n+1)
	// memo[0] = 0
	// memo[1] = 1
	// for i := 2; i < n + 1; i++ {
	// 	memo[i] = memo[i-1] + memo[i-2]
	// }
	// return memo[n]

	ring := []int{0, 1}
	for i := 2; i < n+1; i++ {
		ring[i%2] = ring[0] + ring[1]
	}
	return ring[n%2]
}
