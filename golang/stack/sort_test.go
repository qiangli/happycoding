package stack

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	s := NewStack()
	s.Push(3)
	s.Push(5)
	s.Push(0)
	s.Push(4)
	s.Push(1)
	s.Push(2)

	fmt.Printf("input: %v\n", s)
	sorted := sort(s)
	fmt.Printf("sorted: %v\n", sorted)
	for i := 0; i < 6; i++ {
		v := sorted.Pop()
		if v != i {
			t.FailNow()
		}
	}
}
