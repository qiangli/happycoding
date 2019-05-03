package recursion

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var a []string
	var p func(int, int, int, int)

	s := make([]byte, 2*n)

	p = func(n int, i int, open, close int) {
		fmt.Printf("n: %v i: %v open %v close %v\n", n, i, open, close)

		if close == n {
			fmt.Println("---", string(s))
			a = append(a, string(s))
			return
		}
		if open < n {
			s[i] = '('
			p(n, i+1, open+1, close)
		}
		if open > close {
			s[i] = ')'
			p(n, i+1, open, close+1)
		}
	}

	p(n, 0, 0, 0)

	return a
}
