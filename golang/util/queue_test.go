package util

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	if !q.IsEmpty() {
		t.Fail()
	}

	val := 1

	q.Add(val)

	if q.Size() != 1 {
		t.Fail()
	}

	el := q.Remove()

	if el != val || !q.IsEmpty() {
		t.Fail()
	}
}
