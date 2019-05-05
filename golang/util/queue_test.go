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

	q.Add(val)

	if q.Size() != 1 {
		t.FailNow()
	}

	q.Visit(func(v int) bool {
		fmt.Println(v)
		return true
	})

	el := q.Peek()
	if el != val {
		t.FailNow()
	}

	el = q.Remove()

	if el != val || !q.IsEmpty() {
		t.FailNow()
	}

	size := 100
	for i := 0; i < size; i++ {
		q.Add(i)
	}

	q.Visit(func(v int) bool {
		fmt.Print("->", v)
		return true
	})
	fmt.Println()

	if q.Size() != size {
		t.FailNow()
	}
	for i := 0; i < size; i++ {
		j := q.Remove()
		//fmt.Println(i, j)
		if i != j {
			t.FailNow()
		}
	}

	if q.Size() != 0 {
		t.FailNow()
	}
}
