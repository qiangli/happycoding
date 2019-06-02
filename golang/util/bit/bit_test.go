package bit

import (
	"fmt"
	"testing"
)

func TestBit(t *testing.T) {
	bit := NewBit(300)

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
