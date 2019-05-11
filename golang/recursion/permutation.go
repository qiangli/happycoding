package recursion

import (
	"fmt"
)

func permute(a []int) [][]int {
	if len(a) == 1 {
		return [][]int{a}
	}

	fmt.Println("a", a)

	pick := func(a []int, i int) ([]int, []int) {
		var c []int
		for j, v := range a {
			if j != i {
				c = append(c, v)
			}
		}
		return []int{a[i]}, c
	}

	var r [][]int
	for i := 0; i < len(a); i++ {
		p, c := pick(a, i)
		for _, v := range permute(c) {
			r = append(r, append(p, v...))
		}
	}

	fmt.Println(r)

	return r
}

func permute2(a []int, idx int) {
	if idx == len(a) {
		fmt.Println("permute2", a)
		return
	}

	swap := func(i, j int) {
		a[i], a[j] = a[j], a[i]
	}

	for i := idx; i < len(a); i++ {
		swap(idx, i)
		permute2(a, idx+1)
		swap(idx, i)
	}
}
