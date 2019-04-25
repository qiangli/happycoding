package string

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	s := []byte("abracadabra")
	//abr
	l := 2
	h := hash(s, 0, l)

	fmt.Printf("hash: %v\n %v\n", string(s), h)

	//
	f := rollingHash(s, 0, l)

	for i := 0; i < len(s)-l; i++ {
		hs := hash(s, i, i+l)
		rh := f()
		fmt.Printf("rolling: %v %v\n", rh, hs)
		if hs != rh {
			t.FailNow()
		}
	}
}

func TestNaiveSearch(t *testing.T) {
	s := []byte("abracadabra")
	p := []byte("cad")
	r := NaiveSearch(s, p)
	fmt.Printf("string search:\n %v\n %v\n %v\n", string(s), string(p), r)
}

func TestRabinKarp(t *testing.T) {
	s := []byte("abracadabra")
	p := []byte("cad")
	r := RabinKarp(s, p)
	fmt.Printf("string search:\n %v\n %v\n %v\n", string(s), string(p), r)
}
