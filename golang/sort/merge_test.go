package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := randomArray(15)
	c := mergeSort(a)

	fmt.Printf("input: %v\n", a)
	fmt.Printf("result: %v\n", c)
}

func randomArray(size int) []int {
	c := make([]int, size)
	for i := 0; i < size; i++ {
		c[i] = rand.Intn(100)
	}
	return c
}
