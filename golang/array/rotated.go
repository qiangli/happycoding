package array

import (
	"fmt"
)

//search in sorted and rotated array; return index or -1 if not found
func rotated(a []int, x int, lo, hi int) int {
	if lo == hi {
		if a[lo] == x {
			return lo
		}
		return -1
	}
	mid := (lo + hi) / 2

	fmt.Printf("a: %v x: %v lo: %v hi: %v mid: %v\n", a[lo:hi+1], x, lo, hi, mid)

	if a[lo] <= a[mid] && x <= a[mid] && x >= a[lo] {
		return rotated(a, x, lo, mid)
	} else if a[mid+1] <= a[hi] && x >= a[mid+1] && x <= a[hi] {
		return rotated(a, x, mid+1, hi)
	}

	if a[lo] > a[mid] {
		return rotated(a, x, lo, mid)
	} else {
		return rotated(a, x, mid+1, hi)
	}
}
