package dynamic

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}

	var seg func(string, int) bool
	memo := make(map[int]bool)

	seg = func(a string, i int) bool {
		fmt.Println("idx", i)
		if i == len(a) {
			return true
		}
		if i > len(a) {
			return false
		}

		//
		v, ok := memo[i]
		if ok {
			fmt.Println("cached", i, "v", v)
			return v
		}
		for j := i + 1; j <= len(a); j++ {
			// fmt.Println("i", i, "j", j)
			w := a[i:j]
			found := words[w]
			if found {
				// fmt.Println("found", w, "i", i, "j", j)
				b := seg(a, j)
				fmt.Println("caching", i)
				memo[i] = b
				if b {
					return b
				}
			}
		}
		return false
	}

	return seg(s, 0)
}
