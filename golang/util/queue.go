package util

import (
//"sync"
)

type Queue struct {
	elements []int
}

func NewQueue() *Queue {
	q := &Queue{}
	return q
}

func (r *Queue) Add(v int) {
	r.elements = append(r.elements, v)
}

func (r *Queue) Remove() int {
	el := r.elements[0]
	r.elements = r.elements[1:len(r.elements)]
	return el
}

func (r *Queue) Element() int {
	return r.elements[0]
}

func (r *Queue) IsEmpty() bool {
	return len(r.elements) == 0
}

func (r *Queue) Size() int {
	return len(r.elements)
}
