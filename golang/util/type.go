package util

// Pair defines an int pair
type Pair [2]int

func (r Pair) X() int {
	return r[0]
}

func (r Pair) Y() int {
	return r[1]
}
