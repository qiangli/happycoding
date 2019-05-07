package hash

import (
	"fmt"
	"testing"
)

func TestHashtable(t *testing.T) {
	h := NewHashtable()

	if _, ok := h.Get(1); ok {
		t.FailNow()
	}

	h.Put(1, 3)
	if v, ok := h.Get(1); !ok || v != 3 {
		t.FailNow()
	}

	// string
	h.Put("hello", "world!")
	if v, ok := h.Get("hello"); !ok || v != "world!" {
		t.FailNow()
	}

	fmt.Println(h)

	//
	if b := h.Delete("hello"); !b {
		t.FailNow()
	}
	if _, ok := h.Get("hello"); ok {
		t.FailNow()
	}

	if b := h.Delete(1); !b {
		t.FailNow()
	}

	fmt.Println(h)

	for i := 0; i < 100; i++ {
		h.Put(i, i*1000)
	}

	fmt.Println(h)

	for i := 0; i < 100; i++ {
		if v, ok := h.Get(i); !ok || v != i*1000 {
			t.FailNow()
		}
	}

	for i := 0; i < 100; i++ {
		if ok := h.Delete(i); !ok {
			t.FailNow()
		}
	}

	fmt.Println(h)
}
