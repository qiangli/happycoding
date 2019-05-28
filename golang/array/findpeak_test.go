package array

import (
	"fmt"
	"testing"
)

func TestFindPeak(t *testing.T) {
	a := []int{1, 3, 20, 4, 1, 0}

	r := findPeak(a, 0, len(a)-1)

	fmt.Println("peak: ", r)
}
