package sort

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Merge_sort
// bottomup/iterative
func mergeSort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i = i * 2 {
		fmt.Printf("width: %v\n", i)
		c := []int{}
		for j := 0; j < n; j = j + 2*i {
			l := a[j:min(j+i, n)]
			r := a[min(j+i, n):min(j+2*i, n)]
			fmt.Printf("merging: %v - %v\n", l, r)
			m := merge(l, r)
			c = append(c, m...)
		}
		fmt.Printf("merged: %v\n", c)
		a = c
	}
	return a
}

func merge(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	for i, j, k := 0, 0, 0; k < len(c); k++ {
		if i < len(a) && (j >= len(b) || a[i] < b[j]) {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
