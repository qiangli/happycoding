package bit

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	n := 19
	r := reverse(n)

	fmt.Printf("%v %b %v %b\n", n, n, r, r)
}
