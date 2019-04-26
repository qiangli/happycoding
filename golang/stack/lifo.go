package stack

import (
//"fmt"
)

type node struct {
	data int
	next *node
}

type Stack struct {
	top *node
}

func NewStack() *Stack {
	return &Stack{}
}

func (r *Stack) Push(item int) {
	n := &node{
		data: item,
	}
	n.next = r.top
	r.top = n
}

func (r *Stack) Pop() int {
	n := r.top
	// panic if n is nil
	r.top = n.next
	return n.data
}

func (r Stack) Peek() int {
	// panic if top is nil
	return r.top.data
}

func (r Stack) IsEmpty() bool {
	return r.top == nil
}
