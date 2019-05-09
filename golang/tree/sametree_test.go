package tree

import (
	"fmt"
	"testing"
)

func TestSameTree(t *testing.T) {
	t1 := New(10, 5)
	t2 := New(10, 5)

	fmt.Printf("%v\n%v\n", t1, t2)

	if !Same(t1, t2) {
		t.FailNow()
	}

	t1 = New(10, 5)
	t2 = New(10, 4)

	fmt.Printf("%v\n%v\n", t1, t2)
	if Same(t1, t2) {
		t.FailNow()
	}
}

func TestSameTree2(t *testing.T) {
	t1 := growTree([][3]int{
		{4, 2, 5},
		{2, 1, 3},
		{1, -1, -1},
		{3, -1, -1},
		{5, -1, 6},
		{6, -1, -1},
	})

	t2 := growTree([][3]int{
		{2, 4, 5},
		{4, 1, 3},
		{1, -1, -1},
		{3, -1, -1},
		{5, -1, 6},
		{6, -1, -1},
	})

	if !sametree(t1, t1) {
		t.FailNow()
	}

	if sametree(t1, t2) {
		t.FailNow()
	}
}
