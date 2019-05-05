package util

import ()

type Queue struct {
	head *Node
	tail *Node
	size int
}

// returns false to stop
type Visitor func(int) bool

func NewQueue() *Queue {
	q := &Queue{
		head: &Node{},
		tail: &Node{},
	}
	return q
}

// enqueue
func (r *Queue) Add(v int) {
	n := &Node{
		V: v,
	}
	r.size++

	if r.tail.Next == nil {
		r.head.Next = n
		r.tail.Next = n
		return
	}
	//
	c := r.tail.Next
	r.tail.Next = n
	c.Next = n

}

// dequeue
func (r *Queue) Remove() int {
	r.size--
	c := r.head.Next
	r.head.Next = c.Next
	if r.head.Next == nil {
		r.tail.Next = nil
	}

	return c.V
}

func (r *Queue) Peek() int {
	return r.head.Next.V
}

func (r *Queue) IsEmpty() bool {
	return r.head.Next == nil
}

func (r *Queue) Size() int {
	return r.size
}

func (r *Queue) Visit(f func(int) bool) {
	for c := r.head.Next; c != nil; c = c.Next {
		f(c.V)
	}
}
