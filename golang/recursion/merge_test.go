package recursion

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	size := 12
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Printf("input: %v\n", a)
	b := mergeSort(a)
	fmt.Printf("sorted: %v\n", b)
}
