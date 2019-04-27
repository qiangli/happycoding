package sort

import (
	"math/rand"
)

func randomArray(size int) []int {
	c := make([]int, size)
	for i := 0; i < size; i++ {
		c[i] = rand.Intn(100)
	}
	return c
}
