package string

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm

func NaiveSearch(s, p []byte) int {
	n := len(s)
	m := len(p)
out:
	for i := 0; i < n-m+1; i++ {
		for j := 0; j < m; j++ {
			if s[i+j] != p[j] {
				continue out
			}
		}
		return i
	}
	return -1
}

func hash(k []byte, lo, hi int) int {
	v := 0
	for i := lo; i <= hi; i++ {
		v = ((v%101)*256 + int(k[i])) % 101
	}
	return v
}

func rollingHash(k []byte, lo, hi int) func() int {
	hs := hash(k, lo, hi)
	idx := hi - 1
	m := hi - lo + 1
	return func() int {
		idx++
		// fmt.Printf("rolling idx: %v hash: %v\n", idx, hs)
		if idx == hi {
			return hs
		}
		hs = ((hs+101-int(k[idx-m])*((256%101)*256)%101)*256 + int(k[idx])) % 101
		return hs
	}
}

func RabinKarp(s, p []byte) int {
	n := len(s)
	m := len(p)
	hpat := hash(p, 0, m-1)
	roll := rollingHash(s, 0, m-1)
	//
	match := func(i int) bool {
		fmt.Printf("matching %v\n", i)
		for j := 0; j < m; j++ {
			if s[i+j] != p[j] {
				return false
			}
		}
		return true
	}
	//
	for i := 0; i < n-m+1; i++ {
		rh := roll() //hash(s, i, i + m -1)
		// fmt.Printf("hashing: pattern: % v %v %v\n", hpat, rh, hash(s, i, i + m -1))

		if rh == hpat {
			if match(i) {
				return i
			}
		}
	}
	return 0
}
