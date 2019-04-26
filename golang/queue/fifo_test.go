package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	if !q.IsEmpty() {
		t.FailNow()
	}

	for i := 0; i < 6; i++ {
		q.Add(i)
	}

	if q.IsEmpty() {
		t.FailNow()
	}

	for i := 0; i < 6; i++ {
		v := q.Remove()
		if v != i {
			t.Logf("invalid %v, expecting %v", v, i)
			t.FailNow()
		}
	}

	if !q.IsEmpty() {
		t.FailNow()
	}
}
