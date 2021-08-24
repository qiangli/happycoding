package search

// return index or -1 if not found
func binarySearch(a []int, x int) int {
	lo := 0
	hi := len(a) - 1
	for {

		m := (lo + hi) / 2
		if a[m] == x {
			return m
		} else if x < a[m] {
			hi = m - 1
		} else {
			lo = m + 1
		}
		if lo > hi {
			return -1
		}
	}
}
