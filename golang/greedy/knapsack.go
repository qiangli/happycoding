package greedy

import (
	"fmt"
	"sort"
)

type Item struct {
	value  int
	weight int
}

//https://www.geeksforgeeks.org/fractional-knapsack-problem/
func fractionKnapsack(w int, items []Item) float32 {
	sort.Slice(items, func(i, j int) bool {
		a := float32(items[i].value) / float32(items[i].weight)
		b := float32(items[j].value) / float32(items[j].weight)
		if a < b {
			return true
		}
		return false
	})

	fmt.Println(items)

	var cw int //current weight
	var result float32

	for _, v := range items {
		fmt.Println(cw, result, v)
		remain := w - cw

		if v.weight <= remain {
			cw += v.weight
			result += float32(v.value)
		} else {
			result += float32(remain) * float32(v.value) / float32(v.weight)
			break
		}
	}
	return result
}
