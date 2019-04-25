package array

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	l := NewArrayList()
	l.Append(1)
	l.Append(2)
	//cause resizing
	for i := 4; i < 10; i++ {
		l.Append(i)
	}
	l.Insert(2, 3)

	if l.Get(2) != 3 {
		t.FailNow()
	}

	if l.Size() != 9 {
		t.FailNow()
	}

	fmt.Printf("array list: %v\n", l)
}
