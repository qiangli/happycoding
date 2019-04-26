package string

import (
	"strconv"
)

func compress(s string, pos int) string {
	if pos >= len(s) {
		return ""
	}
	c := s[pos]
	i := pos + 1
	for ; i < len(s); i++ {
		if s[i] != c {
			break
		}
	}
	return string(c) + strconv.Itoa(i-pos) + compress(s, i)
}
