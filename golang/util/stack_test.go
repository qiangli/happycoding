package util

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.FailNow()
	}

	val := 98

	s.Push(val)

	if s.Size() != 1 {
		t.FailNow()
	}

	s.Visit(func(v int) bool {
		fmt.Println(v)
		return true
	})

	el := s.Peek()
	if el != val || s.Size() != 1 {
		t.FailNow()
	}

	el = s.Pop()

	if el != val || !s.IsEmpty() {
		t.FailNow()
	}

	size := 100
	for i := 0; i < size; i++ {
		s.Push(i)
	}

	s.Visit(func(v int) bool {
		fmt.Print("->", v)
		return true
	})
	fmt.Println()

	if s.Size() != size {
		t.FailNow()
	}
	for i := size - 1; i >= 0; i-- {
		j := s.Pop()
		//fmt.Println(i, j)
		if i != j {
			t.FailNow()
		}
	}

	if s.Size() != 0 {
		t.FailNow()
	}
}
