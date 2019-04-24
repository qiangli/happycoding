package dynamic

import (
	"fmt"
	"testing"
)

func TestLcs(t *testing.T) {
	x := []byte{0, 'A', 'B', 'C', 'B', 'D', 'A', 'B'}
	y := []byte{0, 'B', 'D', 'C', 'A', 'B', 'A'}

	//var init func() (int, int, [][]int, [][]int)
	init := func() (int, int, [][]int, [][]int) {
		m := len(x)
		n := len(y)
		c := make([][]int, m)
		b := make([][]int, m)

		for i := 0; i < m; i++ {
			c[i] = make([]int, n)
			b[i] = make([]int, n)
		}
		return m, n, c, b
	}

	m, n, c, b := init()
	lcs(x, y, c, b, m-1, n-1)

	fmt.Printf("lcs: %v %v\n", x, y)
	fmt.Printf("lcs: %v\n%v\n", c, b)

	printLcs(b, x, m-1, n-1)
	fmt.Println()

	//
	fmt.Println("lcs functor")
	m, n, c, b = init()
	f := lcsfunctor(m)
	f(x, y, c, b, m-1, n-1)

	fmt.Printf("lcs functor: %v %v\n", x, y)
	fmt.Printf("lcs functor: %v\n%v\n", c, b)

	printLcs(b, x, m-1, n-1)
	fmt.Println()
}
