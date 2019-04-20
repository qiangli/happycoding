package recursion

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Quicksort
func quickLomuto(a []int, lo, hi int) {
	if lo < hi {
		p := lomuto(a, lo, hi)
		fmt.Printf("partition: %v %v:%v %v\n", p, lo, hi, a)
		quickLomuto(a, lo, p-1)
		quickLomuto(a, p+1, hi)
	}
}

func quickHoare(a []int, lo, hi int) {
	if lo < hi {
		p := hoare(a, lo, hi)
		fmt.Printf("partition: %v %v:%v %v\n", p, lo, hi, a)
		quickHoare(a, lo, p)
		quickHoare(a, p+1, hi)
	}
}

func swap(a []int, i, j int) {
	fmt.Println("swap", i, j)
	a[i], a[j] = a[j], a[i]
}

// Lomuto partition
func lomuto(a []int, lo, hi int) int {
	fmt.Printf("lomuto: %v:%v %v\n", lo, hi, a)
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			swap(a, i, j)
			i++
		}
	}
	swap(a, i, hi)
	return i
}

// Hoare partition
func hoare(a []int, lo, hi int) int {
	pivot := a[(lo+hi)/2]
	i := lo
	j := hi
	for {
		for a[i] < pivot {
			i++
		}
		for a[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		swap(a, i, j)
		i++
		j--
	}
}
