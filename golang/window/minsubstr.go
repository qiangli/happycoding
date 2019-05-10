package window

import (
	"fmt"
)

//https://medium.com/outco/how-to-solve-sliding-window-problems-28d67601a66

func minsubstr(s string, t string) (int, int) {
	counts := make(map[byte]int)
	missing := len(t)
	Inf := 1<<31 - 1
	start, end := 0, Inf

	for i := 0; i < len(t); i++ {
		counts[t[i]] = 0
	}

	slow := 0
	for fast := 0; fast < len(s); fast++ {
		v, ok := counts[s[fast]]
		if ok {
			if v == 0 {
				missing--
			}
			counts[s[fast]]++
		}
		for missing == 0 {
			if fast-slow < end-start {
				start = slow
				end = fast
			}
			_, ok := counts[s[slow]]
			if ok {
				counts[s[slow]]--
				if counts[s[slow]] == 0 {
					missing++
				}
			}
			slow++
		}
		fmt.Println("slow", slow, "fast", fast)
	}

	return start, end
}
