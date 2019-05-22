package recursion

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	head := generate(8)
	print(head)

	r := reverse(head)
	print(r)
}

func TestReversek(t *testing.T) {
	head := generate(8)
	print(head)
	r := reversek(head, 3)
	print(r)

	head = generate(8)
	print(head)
	r = reversek(head, 5)
	print(r)
}

func generate(size int) *node {
	head := &node{
		val: 1,
	}
	p := head
	for i := 1; i < size; i++ {
		n := &node{val: i + 1}
		p.next = n
		p = n
	}
	return head
}

func print(head *node) {
	fmt.Println("")
	for i := head; ; i = i.next {
		fmt.Printf("%v -> ", i.val)
		if i.next == nil {
			fmt.Print("NIL")
			break
		}
	}
	fmt.Println("")
}
