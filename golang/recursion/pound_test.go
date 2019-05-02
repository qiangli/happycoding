package recursion

import (
	"fmt"
	"testing"
)

func TestPondSize(t *testing.T) {
	m := [][]int{
		{0, 2, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
	}

	s := computePondSize(m)

	fmt.Printf("%v\n%v\n", m, s)
}
