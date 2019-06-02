package bit

import (
	"fmt"
)

const Size = 5

type Bit [Size]uint64

func (r *Bit) Set(i int) {
	r[i/64] |= 1 << byte(i%64)
}

func (r *Bit) Get(i int) bool {
	b := r[i/64] & (1 << byte(i%64))
	return b != 0
}

func (r *Bit) Clear(i int) {
	r[i/64] &^= 1 << byte(i%64)
}

func (r *Bit) Toggle(i int) {
	r[i/64] ^= 1 << byte(i%64)
}

func (r *Bit) Negate(n int) Bit {
	var mask Bit
	for i := 0; i < n/64+1; i++ {
		mask[i] = 0xffffffffffffffff
	}
	mask[n/64] >>= byte(64 - n%64)
	bit := r.Clone()
	for i, v := range mask {
		bit[i] = v &^ bit[i]
	}
	return bit
}

func (r *Bit) Clone() Bit {
	var bit Bit
	for i, v := range r {
		bit[i] = v
	}
	return bit
}

func (r Bit) String() string {
	var s string
	for i := len(r); i > 0; i-- {
		s += fmt.Sprintf("%x ", r[i-1])
	}
	return s
}
