package dynamic

import (
	"fmt"
	"testing"
)

func TestCutrod(t *testing.T) {
	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	n := 7

	r := cutrod(p, n)
	fmt.Printf("cutrod: %v %v\n", n, r)

	f := cutrodfunctor()
	fmt.Printf("cutrod functor: %v %v\n", n, f(p, n))
}
