package window

import (
	"fmt"
	"testing"
)

func TestMinsubstr(t *testing.T) {
	s1 := "ADOBECODEBANC"
	s2 := "ABC"

	start, end := minsubstr(s1, s2)
	fmt.Println(start, end)
	if start != 9 || end != 12 {
		t.FailNow()
	}
}
