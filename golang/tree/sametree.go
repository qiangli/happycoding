package tree

import (
	"fmt"
	"math/rand"
)

// https://tour.golang.org/concurrency/8
// https://github.com/golang/tour/blob/master/tree/tree.go
// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(size, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(size) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	visit := func(v *Tree) {
		ch <- v.Value
	}

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	visit(t)

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	//
	same := make(chan bool)

	go func() {
		for {
			c1, ok1 := <-ch1
			c2, ok2 := <-ch2
			if !ok1 || !ok2 {
				same <- ok1 == ok2
				break
			}
			if c1 != c2 {
				same <- false
				break
			}
		}
	}()

	return <-same
}

// https://riptutorial.com/algorithm/example/27566/to-check-if-two-binary-trees-are-same-or-not
func sametree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val &&
		sametree(t1.Left, t1.Left) &&
		sametree(t1.Right, t2.Right)
}
