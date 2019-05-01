package array

import (
	"fmt"
)

// merge b into a where a has large enough buffer at the end to hold b
func merge(a, b []int, la, lb int) {
	i, j, k := la-1, lb-1, la+lb-1
	for ; k >= 0 && j >= 0; k-- {
		fmt.Printf("%v %v %v\n", i, j, k)
		if i >= 0 && a[i] > b[j] {
			a[k] = a[i]
			i--
		} else if j >= 0 {
			a[k] = b[j]
			j--
		}
	}
}
