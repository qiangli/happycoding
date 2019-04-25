package dynamic

import (
	"fmt"
	"testing"
)

func TestActivitySelect(t *testing.T) {
	s := []int{0, 1, 3, 0, 5, 3, 5, 6, 8, 2, 12}
	f := []int{0, 4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}

	a := activitySelect(s, f, 0, len(s)-1)

	fmt.Printf("selection: %v\n", s)
	fmt.Printf("selection: %v\n", f)
	fmt.Printf("selection: %v\n", a)

}
