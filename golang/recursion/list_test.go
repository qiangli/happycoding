package recursion

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	head := generate(16)
	print(head)

	r := reverse(head)
	print(r)
}

func generate(size int) *node {
	head := &node{
		val: 0,
	}
	p := head
	for i := 0; i < size-1; i++ {
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
