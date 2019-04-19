package recursion

// https://en.wikipedia.org/wiki/Linked_list
// node defines linked list node
type node struct {
	val  int
	next *node
}

// reverse the linked list
func reverse(head *node) *node {
	if head == nil {
		return nil
	}

	c := head
	if c.next == nil {
		return c
	}

	r := reverse(c.next)
	c.next.next = c
	c.next = nil

	return r
}
