package util

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	if !q.IsEmpty() {
		t.FailNow()
	}

	val := 98

	q.Enqueue(val)

	if q.Size() != 1 {
		t.FailNow()
	}

	q.Visit(func(v Object) bool {
		fmt.Println(v)
		return true
	})

	el := q.Peek()
	if el != val || q.Size() != 1 {
		t.FailNow()
	}

	el = q.Dequeue()

	if el != val || !q.IsEmpty() {
		t.FailNow()
	}

	size := 100
	for i := 0; i < size; i++ {
		q.Enqueue(i)
	}

	q.Visit(func(v Object) bool {
		fmt.Print("->", v)
		return true
	})
	fmt.Println()

	if q.Size() != size {
		t.FailNow()
	}
	for i := 0; i < size; i++ {
		j := q.Dequeue()
		//fmt.Println(i, j)
		if i != j {
			t.FailNow()
		}
	}

	if q.Size() != 0 {
		t.FailNow()
	}
}

func TestQueueIntArray(t *testing.T) {
	equal := func(a, b Object) bool {
		if len(a.([]int)) != len(b.([]int)) {
			return false
		}
		for i, v := range a.([]int) {
			if b.([]int)[i] != v {
				return false
			}
		}
		return true
	}

	q := NewQueue()

	if !q.IsEmpty() {
		t.FailNow()
	}

	val := []int{12, 18}

	q.Enqueue(val)

	if q.Size() != 1 {
		t.FailNow()
	}

	q.Visit(func(v Object) bool {
		fmt.Println(v)
		return true
	})

	el := q.Peek()
	if !equal(el, val) || q.Size() != 1 {
		t.FailNow()
	}

	el = q.Dequeue()

	if !equal(el, val) || !q.IsEmpty() {
		t.FailNow()
	}

	size := 100
	for i := 0; i < size; i++ {
		q.Enqueue([]int{i, i * 2})
	}

	q.Visit(func(v Object) bool {
		fmt.Print("->", v)
		return true
	})
	fmt.Println()

	if q.Size() != size {
		t.FailNow()
	}
	for i := 0; i < size; i++ {
		j := q.Dequeue()
		fmt.Println(i, j)
		if !equal([]int{i, i * 2}, j) {
			t.FailNow()
		}
	}

	if q.Size() != 0 {
		t.FailNow()
	}
}
