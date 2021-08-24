package queue

type node struct {
	data int
	next *node
}

type Queue struct {
	first *node
	last  *node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (r *Queue) Add(item int) {
	n := &node{
		data: item,
	}
	// init
	if r.first == nil || r.last == nil {
		r.first = n
		r.last = n
		return
	}
	r.last.next = n
	r.last = n
}

func (r *Queue) Remove() int {
	n := r.first
	r.first = n.next
	return n.data
}

func (r Queue) Peek() int {
	return r.first.data
}

func (r Queue) IsEmpty() bool {
	return r.first == nil
}
