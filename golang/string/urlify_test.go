package string

import (
	"fmt"
	"testing"
)

func TestURLify(t *testing.T) {
	a := []byte("Mr John Smith----")
	l := 13

	fmt.Printf("urlify: %v %v\n", string(a), l)

	URLify(a, l)

	fmt.Printf("urlify: %v %v\n", string(a), l)
}
