package recursion

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println(a)
	p := permute(a)
	fmt.Println("permute", p)
}

func TestPermute2(t *testing.T) {
	tests := [][]int{
		[]int{0},
		[]int{0, 1},
		[]int{0, 1, 2},
		[]int{0, 1, 2, 3},
		[]int{0, 1, 2, 3, 4},
	}
	for _, tc := range tests {
		p := permute2(tc, 0)
		t.Logf("len: %v %v", len(p), p)
	}
}
