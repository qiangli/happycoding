package sort

func bubbleSort(a []int) {
	swap := func(a []int, i, j int) {
		a[i], a[j] = a[j], a[i]
	}
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				swap(a, j, j+1)
			}
		}
	}
}
