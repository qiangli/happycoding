package array

import (
	"fmt"
	"sort"
)

//https://www.geeksforgeeks.org/suffix-array-set-1-introduction/

type SuffixArray struct {
	s string
	a []int
}

func (r SuffixArray) Len() int {
	return len(r.a)
}

func (r SuffixArray) Swap(i, j int) {
	r.a[i], r.a[j] = r.a[j], r.a[i]
}

func (r SuffixArray) Less(i, j int) bool {
	sz := len(r.a)
	s1, s2 := r.a[i], r.a[j]
	for k := 0; s1+k < sz && s2+k < sz; k++ {
		c1, c2 := r.s[s1+k], r.s[s2+k]
		if c1 > c2 {
			return false
		}
		if c1 < c2 {
			return true
		}
	}
	return i > j
}

// binary search
func (r SuffixArray) Search(p string) int {
	n := len(r.a)
	m := len(p)
	match := func(k int) int {
		x := r.a[k]
		i := 0
		for ; i < m && x+i < n; i++ {
			c := r.s[x+i]
			if p[i] > c {
				return 1
			} else if p[i] < c {
				return -1
			}
		}

		if m <= n-x {
			return 0
		}
		//pattern is longer than suffix
		return 1
	}

	lo, hi := 0, len(r.a)-1
	for {
		if lo > hi {
			break
		}
		mid := lo + (hi-lo)/2
		x := match(mid)
		if x == 0 {
			return mid
		}
		if x < 0 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func (r SuffixArray) String() string {
	var s []string
	for _, v := range r.a {
		s = append(s, r.s[v:])
	}
	return fmt.Sprintf("%v\n", s)
}

func NewSuffixArray(s string) *SuffixArray {
	a := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		a[i] = i
	}
	sa := SuffixArray{
		s: s,
		a: a,
	}
	sort.Sort(sa)

	return &sa
}
