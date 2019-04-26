package list

import (
	"fmt"
)

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (r *DoublyLinkedList) Append(v int) {
	n := &Node{
		data: v,
	}

	//init
	if r.head == nil || r.tail == nil {
		r.head = n
		r.tail = n
		return
	}

	n.prev = r.tail
	n.prev.next = n
	r.tail = n
}

func (r *DoublyLinkedList) Iterate(f func(*Node)) {
	c := r.head
	for {
		f(c)
		if c.next == nil {
			break
		}
		c = c.next
	}
}

func Convert2Tree(n, c *Node) {
	fmt.Printf("converting: %v c: %v\n", n, c)
	if n == nil {
		return
	}
	nn := n.next

	n.prev = c
	n.next = nil

	var cc *Node
	if c != nil {
		n.next = c.next
		if c.next != nil {
			cc = c.next.next
		}
	}
	Convert2Tree(nn, cc)
	fmt.Printf("converted: %v children: %v %v\n", n, n.prev, n.next)
}

type queue struct {
	items []*Node
}

func NewQueue() *queue {
	return &queue{}
}

func (r *queue) enqueue(v *Node) {
	r.items = append(r.items, v)
	// fmt.Printf("enqueue %v\n", r.items)
}

func (r *queue) dequeue() *Node {
	n := r.items[0]
	r.items = r.items[1:]
	// fmt.Printf("dequeue %v\n", r.items)
	return n
}

func (r *queue) empty() bool {
	return len(r.items) == 0
}

func BFS(root *Node) {
	fmt.Printf("queue: root %v\n", root)

	q := NewQueue()
	q.enqueue(root)

	visit := func(n *Node) {
		fmt.Printf("visit: %v\n", n)
	}

	for i := 0; !q.empty() && i < 6; i++ {
		n := q.dequeue()

		visit(n)

		if n.prev != nil {
			q.enqueue(n.prev)
		}
		if n.next != nil {
			q.enqueue(n.next)
		}
	}
}
