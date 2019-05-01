package array

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	a := []int{2, 3, 6, 7, 10, -1, -1, -1, -1}
	b := []int{1, 4, 5, 8}

	fmt.Printf("%v\n%v\n", a, b)
	merge(a, b, 5, 4)
	fmt.Printf("%v\n", a)
}
