package recursion

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 6, 8, 9, 10, 20, 30, 50}
	x := 5
	r := binarySearch(a, x, 0, len(a)-1)

	fmt.Printf("binary search: %v\n %v %v\n", a, x, r)
}
