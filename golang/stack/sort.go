package stack

import (
	"fmt"
)

func sort(s *Stack) *Stack {
	fmt.Printf("sort stack: %v\n", s)
	// sorted if empty or single item
	if s.IsEmpty() {
		return s
	}
	x := s.Pop()
	if s.IsEmpty() {
		s.Push(x)
		return s
	}
	//
	s.Push(x)
	// split s into two stacks and sort
	fmt.Printf("sorting: %v\n", s)

	s1 := NewStack()
	s2 := NewStack()
	for {
		a := s.Pop()
		s1.Push(a)
		if s.IsEmpty() {
			break
		}
		b := s.Pop()
		s2.Push(b)
		if s.IsEmpty() {
			break
		}
	}
	fmt.Printf("split: %v\n%v\n", s1, s2)

	s1 = sort(s1)
	s2 = sort(s2)
	// merge
	fmt.Printf("merge: %v\n%v\n", s1, s2)

	ns := NewStack()
	drain := func(from *Stack, to *Stack) {
		for {
			if from.IsEmpty() {
				break
			}
			to.Push(from.Pop())
		}
	}
	pump := func(from *Stack, to *Stack, threshold int) {
		for {
			if from.IsEmpty() {
				break
			}
			i := from.Peek()
			if i <= threshold {
				to.Push(from.Pop())
			} else {
				break
			}
		}
	}
	for {
		if s1.IsEmpty() && s2.IsEmpty() {
			break
		}
		if s1.IsEmpty() {
			drain(s2, ns)
			break
		} else if s2.IsEmpty() {
			drain(s1, ns)
			break
		}
		a := s1.Peek()
		b := s2.Peek()
		if a < b {
			pump(s1, ns, b)
		} else {
			pump(s2, ns, a)
		}
	}
	fmt.Printf("merged: %v\n", ns)

	//reverse
	rs := NewStack()
	for {
		if ns.IsEmpty() {
			break
		}
		x := ns.Pop()
		rs.Push(x)
	}
	fmt.Printf("sorted: %v\n", rs)
	return rs
}
