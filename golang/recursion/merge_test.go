package recursion

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := randomArray(11)
	fmt.Printf("input: %v\n", a)
	b := mergeSort(a)
	fmt.Printf("sorted: %v\n", b)
}

func randomArray(size int) []int {
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = rand.Intn(100)
	}
	return a
}
