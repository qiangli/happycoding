package bit

import (
	"fmt"
	"testing"
)

func TestBit(t *testing.T) {
	var bit Bit

	for i := 0; i <= 300; i++ {
		if bit.Get(i) {
			t.FailNow()
		}
		if bit.Set(i); !bit.Get(i) {
			t.FailNow()
		}
		if bit.Clear(i); bit.Get(i) {
			t.FailNow()
		}
		if bit.Toggle(i); !bit.Get(i) {
			t.FailNow()
		}

		if bit.Toggle(i); bit.Get(i) {
			t.FailNow()
		}
		fmt.Println(bit)
	}
}

func TestBitClone(t *testing.T) {
	var bit Bit
	bit.Set(1)
	bit.Set(7)

	fmt.Println(bit)

	cp := bit.Clone()
	if !cp.Get(1) || !cp.Get(7) {
		t.FailNow()
	}
	fmt.Println(cp)
}

func TestBitNegate(t *testing.T) {
	var bit Bit
	bit.Set(0)
	bit.Set(1)
	bit.Set(2)

	fmt.Println(bit)

	ne := bit.Negate(3)
	fmt.Println(ne)
}
