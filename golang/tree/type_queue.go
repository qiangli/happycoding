package tree

// type Object interface{}

// // returns false to stop
// type Visitor func(Object) bool

// type node struct {
// 	val  Object
// 	next *node
// }

type Queue struct {
	head *node
	tail *node
	size int
}

func NewQueue() *Queue {
	q := &Queue{
		head: &node{},
		tail: &node{},
	}
	return q
}

// Enqueue adds item to the end
func (r *Queue) Enqueue(v Object) {
	n := &node{
		val: v,
	}
	r.size++

	if r.tail.next == nil {
		r.head.next = n
		r.tail.next = n
		return
	}
	//
	c := r.tail.next
	r.tail.next = n
	c.next = n

}

// Dequeue removes item from head
func (r *Queue) Dequeue() Object {
	r.size--
	c := r.head.next
	r.head.next = c.next
	if r.head.next == nil {
		r.tail.next = nil
	}

	return c.val
}

func (r *Queue) IsEmpty() bool {
	return r.head.next == nil
}

func (r *Queue) Peek() Object {
	return r.head.next.val
}

func (r *Queue) Size() int {
	return r.size
}

func (r *Queue) Visit(f Visitor) {
	for c := r.head.next; c != nil; c = c.next {
		if !f(c.val) {
			break
		}
	}
}
