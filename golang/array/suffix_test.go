package array

import (
	"fmt"
	"testing"
)

func TestNewSuffixArray(t *testing.T) {
	s := "banana"
	fmt.Println(s)

	sa := NewSuffixArray(s)

	fmt.Println(sa)

	idx := sa.Search("ana")

	fmt.Println("search", idx)
}
