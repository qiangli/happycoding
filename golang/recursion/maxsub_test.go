package recursion

import (
	"fmt"
	"testing"
)

func TestMaxsub(t *testing.T) {
	a := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	fmt.Printf("maxsub: %v\n", a)

	l, r, max := maxsub(a, 0, len(a)-1)
	fmt.Printf("maxsub: %v, %v max: %v\n", l, r, max)
}
