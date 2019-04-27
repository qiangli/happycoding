package bit

import (
	"fmt"
	"testing"
)

func TestBit(t *testing.T) {
	n := 0x00ff
	b := get(n, 10)
	fmt.Printf("bit: %v %v\n", n, b)

	n = set(n, 10)
	b = get(n, 10)
	fmt.Printf("bit: %v %v\n", n, b)

	n = clear(n, 10)
	b = get(n, 10)
	fmt.Printf("bit: %v %v\n", n, b)
}

func TestBitOp(t *testing.T) {
	var n int
	var s string
	//
	n = 0x7fffffffffffffff
	s = fmt.Sprintf("%b", int64(n))
	fmt.Printf("bit:   %v %v\n", s, len(s))

	n = 1 << 62
	s = fmt.Sprintf("%b", int64(n))
	fmt.Printf("bit:   %v %v\n", s, len(s))

	n = -1 << 63
	s = fmt.Sprintf("%b", int64(n))
	fmt.Printf("bit: %v %d %v\n", s, n, len(s))

	n = int(0x7fffffffffffffff)
	s = fmt.Sprintf("%b", int64(n))
	fmt.Printf("bit: %v %d %v\n", s, n, len(s))
}
