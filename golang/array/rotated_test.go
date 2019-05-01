package array

import (
	"fmt"
	"testing"
)

func TestRotated(t *testing.T) {
	a := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}

	fmt.Println(a)
	if rotated(a, 20, 0, len(a)-1) != 3 {
		t.FailNow()
	}

	if rotated(a, 3, 0, len(a)-1) != 6 {
		t.FailNow()
	}

	if rotated(a, 100, 0, len(a)-1) != -1 {
		t.FailNow()
	}
}
