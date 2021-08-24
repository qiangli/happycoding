package string

//"fmt"

func reverse(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		s[i], s[j] = s[j], s[i]
	}
}
