package stack

import (
	"fmt"
	"testing"
)

func TestLIFO(t *testing.T) {
	s := NewStack()
	for i := 0; i < 6; i++ {
		s.Push(i)
	}

	fmt.Printf("stack: %v\n", s)

	if s.IsEmpty() {
		t.Log("empty stack")
		t.FailNow()
	}

	for i := 5; i >= 0; i-- {
		v := s.Pop()
		t.Logf("popping val: %v, expecting: %v", v, i)

		if v != i {
			t.FailNow()
		}
	}

	fmt.Printf("stack: %v\n", s)

	if !s.IsEmpty() {
		t.FailNow()
	}
}
