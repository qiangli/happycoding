package string

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	s := "aabcccccaaa"

	c := compress(s, 0)

	fmt.Printf("compress: %v\n%v\n", s, c)
}
