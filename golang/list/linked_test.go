package list

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	for i := 1; i < 7; i++ {
		q.enqueue(&Node{data: i})
	}
	fmt.Printf("queue %v\n", q)
	for i := 1; i < 7; i++ {
		q.dequeue()
	}
	if !q.empty() {
		t.FailNow()
	}
}

func TestDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList()
	for i := 1; i < 7; i++ {
		list.Append(i)
	}

	list.Iterate(func(n *Node) {
		fmt.Printf("%v\n", n.data)
	})

	root := list.head
	fmt.Printf("%v\n", root)

	Convert2Tree(root, root.next)
	BFS(root)
}
