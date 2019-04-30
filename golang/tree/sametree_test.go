package tree

import (
	"fmt"
	"testing"
)

func TestSameTree(t *testing.T) {
	t1 := New(10, 5)
	t2 := New(10, 5)

	fmt.Printf("%v\n%v\n", t1, t2)
	fmt.Println(Same(t1, t2))

	t1 = New(10, 5)
	t2 = New(10, 4)

	fmt.Printf("%v\n%v\n", t1, t2)
	fmt.Println(Same(t1, t2))
}
