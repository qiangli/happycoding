package recursion

import (
	"fmt"
)

// find maximum subarray
// introduction to algorithms - 3rd edition p72

// find max crossing subarray
func maxxsub(a []int, lo, mid, hi int) (left, right, max int) {
	left = mid
	lsum := a[left]
	sum := 0
	for i := mid; i > lo; i-- {
		sum += a[i]
		if sum > lsum {
			lsum = sum
			left = i
		}
	}
	right = mid + 1
	rsum := a[right]
	sum = 0
	for j := mid + 1; j <= hi; j++ {
		sum += a[j]
		if sum > rsum {
			rsum = sum
			right = j
		}
	}
	max = lsum + rsum
	// fmt.Printf("cross: %v %v %v %v\n", lo, mid, hi, a)
	// fmt.Printf("cross: %v %v max: %v\n",left, right, max)
	return
}

// find max subarray
func maxsub(a []int, lo, hi int) (int, int, int) {
	if lo == hi {
		return lo, hi, a[lo]
	}
	mid := (lo + hi) / 2

	fmt.Printf("lo: %v hi: %v mid: %v\n", lo, hi, mid)

	llo, lhi, lmax := maxsub(a, lo, mid)
	rlo, rhi, rmax := maxsub(a, mid+1, hi)
	xlo, xhi, xmax := maxxsub(a, lo, mid, hi)

	if lmax >= rmax && lmax >= xmax {
		return llo, lhi, lmax
	}
	if rmax >= lmax && rmax >= xmax {
		return rlo, rhi, rmax
	}
	return xlo, xhi, xmax
}
