package string

import (
	"fmt"
)

func anagram(s1, s2 string) bool {
	cnt := make(map[rune]int)
	for _, c := range s1 {
		v := cnt[c]
		cnt[c] = v + 1
	}

	fmt.Println(cnt)

	for _, c := range s2 {
		v := cnt[c]
		cnt[c] = v - 1
	}

	fmt.Println(cnt)

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
