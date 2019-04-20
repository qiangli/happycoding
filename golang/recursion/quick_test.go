package recursion

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLomuto(t *testing.T) {
	a := randomData(10)
	fmt.Printf("input: %v\n", a)

	quickLomuto(a, 0, len(a)-1)
	fmt.Printf("sorted: %v\n", a)
}
func TestHoare(t *testing.T) {
	a := randomData(15)
	fmt.Printf("input: %v\n", a)

	quickHoare(a, 0, len(a)-1)
	fmt.Printf("sorted: %v\n", a)
}

func randomData(size int) []int {
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = rand.Intn(100)
	}
	return a
}
