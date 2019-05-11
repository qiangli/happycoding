package string

import (
	"fmt"
	"testing"
)

func TestPermutaion(t *testing.T) {
	a1 := "god"
	a2 := "dog"
	b1, b2 := permutation(a1, a2), permutation2(a1, a2)
	fmt.Printf("permutation %v %v %v %v\n", a1, a2, b1, b2)
	if !b1 || !b2 {
		t.FailNow()
	}
}

func TestPermutaion3(t *testing.T) {
	s := "ABCA"
	p := permutation3(s)
	fmt.Println(p, len(p))
}
