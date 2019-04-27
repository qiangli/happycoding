package recursion

import ()

// return -1 if not found
func binarySearch(a []int, x int, lo, hi int) int {
	if lo > hi {
		return -1
	}
	m := (lo + hi) / 2
	if a[m] == x {
		return m
	}

	if x < a[m] {
		return binarySearch(a, x, lo, m-1)
	}
	return binarySearch(a, x, m+1, hi)
}
