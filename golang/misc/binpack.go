package misc

import (
	"fmt"
	"sort"
)

func binpack(weight []int, capacity int) {
	sort.Ints(weight)

	fmt.Println(weight)

	cnt := 0
	bins := make([]int, len(weight))

	fit := func() {
		for i, v := range weight {
			j := 0
			for ; j < cnt; j++ {
				if v <= capacity-bins[j] {
					bins[j] += v
					fmt.Printf("pack item %v %v in bin %v\n", i, v, j)
					break
				}
			}
			if j == cnt {
				bins[cnt] = v
				fmt.Printf("pack item %v %v in new bin %v\n", i, v, cnt)
				cnt++
			}
		}
	}

	fit()

	fmt.Println(bins, cnt)
}
