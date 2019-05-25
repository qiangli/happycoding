package string

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	s := []byte("abcdefgh")
	fmt.Printf("input: %v\n", string(s))
	reverse(s)
	fmt.Printf("reverse: %v\n", string(s))
}
