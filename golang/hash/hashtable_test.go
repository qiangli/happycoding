package hash

import (
	"testing"
)

func TestHashtable(t *testing.T) {
	h := NewHashtable()

	if v, ok := h.Get(1); ok || v == 3 {
		t.FailNow()
	}

	h.Put(1, 3)

	if v, ok := h.Get(1); !ok || v != 3 {
		t.FailNow()
	}
}
