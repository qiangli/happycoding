package string

import (
	"fmt"
	"sort"
)

type sortable []rune

func (r sortable) Len() int {
	return len(r)
}
func (r sortable) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r sortable) Less(i, j int) bool {
	return r[i]-r[j] < 0
}

func (r sortable) Equal(s sortable) bool {
	if s == nil || len(r) != len(s) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if r[i] != s[i] {
			return false
		}
	}
	return true
}

func permutation(a1 string, a2 string) bool {
	if len(a1) != len(a2) {
		return false
	}
	s1 := sortable(a1)
	s2 := sortable(a2)
	sort.Sort(s1)
	sort.Sort(s2)

	fmt.Printf("permutation: %v %v\n", s1, s2)
	return s1.Equal(s2)
}

// check permutation
func permutation2(a1, a2 string) bool {
	if len(a1) != len(a2) {
		return false
	}

	s1 := []rune(a1)
	s2 := []rune(a2)

	charset := make([]rune, 0xffff) //assumption - increase to deal with large char set

	// increment or decrement depending on whether the char is in s1 or s2
	for i := 0; i < len(s1); i++ {
		charset[s1[i]]++
		charset[s2[i]]--
	}

	for i := 0; i < len(s1); i++ {
		if charset[i] != 0 {
			return false
		}
	}
	return true
}

// https://www.geeksforgeeks.org/distinct-permutations-string-set-2/
// print all permutations
func permutation3(s string) []string {
	sz := len(s)

	swap := func(a []byte, i, j int) {
		a[i], a[j] = a[j], a[i]
	}
	duped := func(a []byte, idx, i int) bool {
		for j := idx; j < i; j++ {
			if a[j] == a[i] {
				return true
			}
		}
		return false
	}

	var res []string
	var permute func([]byte, int)
	permute = func(a []byte, idx int) {
		if idx == sz {
			fmt.Println(a)
			res = append(res, string(a))
		}

		for i := idx; i < sz; i++ {
			if duped(a, idx, i) {
				continue
			}
			swap(a, idx, i)
			permute(a, idx+1)
			swap(a, idx, i)
		}
	}

	permute([]byte(s), 0)
	return res
}
