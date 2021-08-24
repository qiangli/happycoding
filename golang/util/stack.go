package util

type Stack struct {
	head *Node
	size int
}

func NewStack() *Stack {
	return &Stack{
		head: &Node{},
	}
}

func (r *Stack) Push(v int) {
	n := &Node{
		V: v,
	}
	r.size++

	if r.head == nil {
		r.head.Next = n
		return
	}
	//
	c := r.head.Next
	r.head.Next = n
	n.Next = c
}

func (r *Stack) Pop() int {
	c := r.head.Next
	r.size--
	r.head.Next = c.Next
	return c.V
}

func (r *Stack) IsEmpty() bool {
	return r.head.Next == nil
}

func (r *Stack) Peek() int {
	return r.head.Next.V
}

func (r *Stack) Size() int {
	return r.size
}

func (r *Stack) Visit(f Visitor) {
	for c := r.head.Next; c != nil; c = c.Next {
		if !f(c.V) {
			break
		}
	}
}
