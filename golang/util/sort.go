package util

// https://www.codementor.io/tucnak/using-golang-for-competitive-programming-h8lhvxzt3
type pair struct {
	x int
	y float64
}

type pairs []pair                  /**/
func (r pairs) Len() int           { return len(r) }
func (r pairs) Less(i, j int) bool { return r[i].x < r[j].x || r[i].y < r[j].y }
func (r pairs) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
