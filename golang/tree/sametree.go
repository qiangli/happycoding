//https://tour.golang.org/concurrency/8

package tree

import (
	"fmt"
	"sync"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	visit := func(v *Tree) {
		//fmt.Printf(" -> %v", v.Value)
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

	collect := func(ch chan int) []int {
		var a []int
		for v := range ch {
			a = append(a, v)
		}
		return a
	}

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
	var a1, a2 []int
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		a1 = collect(ch1)
		wg.Done()
	}()

	go func() {
		a2 = collect(ch2)
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("%v\n%v\n", a1, a2)

	if len(a1) != len(a2) {
		return false
	}
	for i, v := range a1 {
		if a2[i] != v {
			return false
		}
	}
	return true
}
