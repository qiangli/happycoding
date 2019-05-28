package array

import (
	"fmt"
)

// https://www.geeksforgeeks.org/find-a-peak-in-a-given-array/
func findPeak(a []int, lo, hi int) int {
	fmt.Println(lo, hi)

	n := len(a)
	mid := lo + (hi-lo)/2
	if (mid == 0 || a[mid-1] <= a[mid]) && (mid == n-1 || a[mid] >= a[mid+1]) {
		return mid
	}

	if mid > 0 && a[mid-1] > a[mid] {
		return findPeak(a, lo, mid-1)
	}
	return findPeak(a, mid+1, hi)
}
