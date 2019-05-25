package string

//https://www.geeksforgeeks.org/program-count-vowels-string-iterative-recursive/
func countVowel(s []byte, n int) int {
	cnt := func(b byte) int {
		for _, v := range []byte("aeiou") {
			if v == b {
				return 1
			}
		}
		return 0
	}
	//println(n)

	if n == 0 {
		return 0
	}
	return countVowel(s, n-1) + cnt(s[n-1])
}
