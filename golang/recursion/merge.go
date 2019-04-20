package recursion

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Merge_sort
// topdown/recursion
func mergeSort(a []int) []int {
	m := len(a) / 2
	if m == 0 {
		return a
	}
	l := a[0:m]
	r := a[m:len(a)]
	fmt.Printf("%v - %v\n", l, r)

	l = mergeSort(l)
	r = mergeSort(r)

	fmt.Printf("sorted: %v - %v\n", l, r)
	return merge(l, r)
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
