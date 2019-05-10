package bit

func reverse(n int) int {
	rev := 0
	for n > 0 {
		rev <<= 1
		if n&1 == 1 {
			rev ^= 1
		}
		n >>= 1
	}
	return rev
}
