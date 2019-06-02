package bit

import (
	"fmt"
)

type Bit []uint64

func NewBit(size int) Bit {
	bit := make([]uint64, size/64+1)
	return bit
}

func (r Bit) Set(i int) {
	r[i/64] |= 1 << byte(i%64)
}

func (r Bit) Get(i int) bool {
	return r[i/64]&(1<<byte(i%64)) != 0
}

func (r Bit) Clear(i int) {
	r[i/64] &^= 1 << byte(i%64)
}

func (r Bit) Toggle(i int) {
	r[i/64] ^= 1 << byte(i%64)
}

func (r Bit) String() string {
	var s string
	for i := len(r); i > 0; i-- {
		s += fmt.Sprintf("%x ", r[i-1])
	}
	return s
}
