package misc

import (
	"fmt"
	"sort"
)

//https://www.geeksforgeeks.org/combinations-with-repetitions/

// https://www.geeksforgeeks.org/print-all-permutations-with-repetition-of-characters/
func permutationChar(s string) {
	var n = len(s)
	var ba = make([]byte, n)
	for i := 0; i < n; i++ {
		ba[i] = s[i]
	}

	var permutate func([]byte, int)
	permutate = func(data []byte, idx int) {
		if idx == n {
			//fmt.Println(data)
			for _, v := range data {
				fmt.Printf("%c", v)
			}
			fmt.Println()
			return
		}
		for i := 0; i < n; i++ {
			data[idx] = ba[i]
			permutate(data, idx+1)
		}
	}

	sort.Slice(ba, func(i, j int) bool {
		if ba[i] < ba[j] {
			return true
		}
		return false
	})
	data := make([]byte, n)
	permutate(data, 0)
}

// https://www.geeksforgeeks.org/print-all-possible-combinations-of-r-elements-in-a-given-array-of-size-n/
func combination(a []int, N, K int) {
	var nk func([]int, int, int)
	nk = func(data []int, n, k int) {
		if k == K {
			fmt.Println(data)
			return
		}
		if n == N {
			return
		}

		data[k] = a[n]
		nk(data, n+1, k+1)
		nk(data, n+1, k)
	}

	data := make([]int, K)
	nk(data, 0, 0)
}

// https://www.geeksforgeeks.org/make-combinations-size-k/
func combinationNK(N, K int) {
	var nk func([]int, int, int)
	nk = func(data []int, start, idx int) {
		//fmt.Println("nk start/idx", start, idx)
		if idx == K {
			fmt.Println(data)
			return
		}

		for i := start; i <= N; i++ {
			//
			if N-i+1 < K-idx {
				continue
			}
			data[idx] = i
			nk(data, i+1, idx+1)
		}
	}

	data := make([]int, K)
	nk(data, 1, 0)
}

func combinationNK2(N, K int) {
	var nk func([]int, int, int)
	nk = func(data []int, n, k int) {
		if k == K {
			fmt.Println(data)
			return
		}
		if n == N {
			return
		}

		data[k] = n + 1

		nk(data, n+1, k+1)
		nk(data, n+1, k)
	}

	data := make([]int, K)
	nk(data, 0, 0)
}

// https://www.geeksforgeeks.org/print-all-combinations-of-points-that-can-compose-a-given-number/
func compose(a []int, n, K, i int) {
	//fmt.Println(n, i, a, a[:i])
	switch {
	case n == 0:
		//fmt.Println("output", n, i, a)
		fmt.Println(a[:i])
	case n > 0:
		for k := 1; k <= K; k++ {
			a[i] = k
			compose(a, n-k, K, i+1)
		}
	default:
		//fmt.Println("n", n)
	}
}

// https://www.geeksforgeeks.org/find-all-combinations-that-adds-upto-given-number-2/
func compose2(a []int, n, K, i int) {
	switch {
	case n == 0:
		fmt.Println(a[:i])
	case n > 0:
		start := 1
		if i > 0 {
			start = a[i-1]
		}
		//
		for k := start; k <= K; k++ {
			a[i] = k
			compose2(a, n-k, K, i+1)
		}
	default:
		//fmt.Println("n", n)
	}
}
