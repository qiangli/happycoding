package util

type Node struct {
	V    int
	Prev *Node
	Next *Node
}

// returns false to stop
type Visitor func(int) bool
