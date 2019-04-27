package sort

import ()

func selectionSort(a []int) {
	swap := func(a []int, i, j int) {
		a[i], a[j] = a[j], a[i]
	}
	n := len(a)
	for i := 0; i < n; i++ {
		min := i
		for j := i; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		swap(a, min, i)
	}
}
