package util

import ()

type Object interface{}

type Stack struct {
	head *node
	size int
}

type node struct {
	val  Object
	next *node
}

func NewStack() *Stack {
	return &Stack{
		head: &node{},
	}
}

func (r *Stack) Push(v int) {
	n := &node{
		val: v,
	}
	r.size++

	if r.head == nil {
		r.head.next = n
		return
	}
	//
	c := r.head.next
	r.head.next = n
	n.next = c
}

func (r *Stack) Pop() Object {
	c := r.head.next
	r.size--
	r.head.next = c.next
	return c.val
}

func (r *Stack) IsEmpty() bool {
	return r.head.next == nil
}

func (r *Stack) Peek() Object {
	return r.head.next.val
}

func (r *Stack) Size() int {
	return r.size
}

// returns false to stop
type Visitor func(Object) bool

func (r *Stack) Visit(f Visitor) {
	for c := r.head.next; c != nil; c = c.next {
		if !f(c.val) {
			break
		}
	}
}
