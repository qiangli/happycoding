package string

import (
	"fmt"
)

// https://www.geeksforgeeks.org/recursive-function-to-do-substring-search/
func contain(text string, pat string) bool {
	fmt.Println(text, pat)

	t := []byte(text)
	p := []byte(pat)

	var sub func(int, int) bool
	sub = func(i, j int) bool {
		//fmt.Println(i, j)
		if j >= len(p) {
			return true
		}

		if i >= len(t) {
			return false
		}

		if t[i] == p[j] {
			return sub(i+1, j+1)
		}

		return sub(i+1, 0)
	}

	return sub(0, 0)
}
