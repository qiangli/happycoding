package recursion

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Linked_list
// node defines linked list node
type node struct {
	val  int
	next *node
}

// reverse the linked list
func reverse(c *node) *node {
	fmt.Println("reversing c", c)

	if c == nil || c.next == nil {
		return c
	}

	tail := reverse(c.next)
	c.next.next = c
	c.next = nil
	fmt.Println("c", c.val, "tail", tail)
	return tail
}

// https://www.geeksforgeeks.org/reverse-a-list-in-groups-of-given-size/
func reversek(c *node, k int) *node {
	if c == nil {
		return nil
	}

	head := c
	tail := c
	for i := 1; i < k; i++ {
		if tail.next != nil {
			tail = tail.next
		}
	}
	next := tail.next
	tail.next = nil

	r := reverse(head)
	if next == nil {
		return r
	}

	head.next = reversek(next, k)
	return r
}
