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

	lo, hi := 0, len(r.a)
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

// https://en.wikipedia.org/wiki/Suffix_array
func (r SuffixArray) Match(p string) (int, int) {
	n := len(r.a)
	m := len(p)
	match := func(k int) (bool, int) {
		x := r.a[k]
		i := 0
		for ; i < m && x+i < n; i++ {
			c := r.s[x+i]
			if p[i] > c {
				return false, 1
			} else if p[i] < c {
				return false, -1
			}
		}
		d := m - (n - x)
		return d <= 0, d
	}

	start := func(lo, hi int) int {
		idx := -1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			b, x := match(mid)
			if b {
				idx = mid
			}
			if b || x < 0 {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		return idx
	}
	end := func(lo, hi int) int {
		idx := -1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			b, x := match(mid)
			if b {
				idx = mid
			}
			if b || x > 0 {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return idx
	}
	lo, hi := 0, len(r.a)-1
	s := start(lo, hi)
	if s == -1 {
		return -1, -1
	}
	e := end(s, hi)
	return s, e
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
