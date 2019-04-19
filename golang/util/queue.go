package util

import (
	"sync"
)

type Queue struct {
	elements []int
	sync.RWMutex
}

func NewQueue() *Queue {
	q := &Queue{}
	return q
}

func (r *Queue) Add(v int) {
	r.Lock()
	defer r.Unlock()
	r.elements = append(r.elements, v)
}

func (r *Queue) Remove() int {
	r.Lock()
	defer r.Unlock()
	el := r.elements[0]
	r.elements = r.elements[1:len(r.elements)]
	return el
}

func (r *Queue) Element() int {
	r.RLock()
	defer r.RUnlock()
	return r.elements[0]
}

func (r *Queue) IsEmpty() bool {
	r.RLock()
	defer r.RUnlock()
	return len(r.elements) == 0
}

func (r *Queue) Size() int {
	r.RLock()
	defer r.RUnlock()
	return len(r.elements)
}
