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
