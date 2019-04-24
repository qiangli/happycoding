package dynamic

import (
	"fmt"
	"testing"
)

func TestLcs(t *testing.T) {
	x := []byte{0, 'A', 'B', 'C', 'B', 'D', 'A', 'B'}
	y := []byte{0, 'B', 'D', 'C', 'A', 'B', 'A'}

	m := len(x)
	n := len(y)
	c := make([][]int, m)
	b := make([][]int, m)

	for i := 0; i < m; i++ {
		c[i] = make([]int, n)
		b[i] = make([]int, n)
	}

	lcs(x, y, c, b, m-1, n-1)

	fmt.Printf("lcs: %v %v\n", x, y)
	fmt.Printf("lcs: %v\n", c)

	printLcs(b, x, m-1, n-1)
	fmt.Println()
}
