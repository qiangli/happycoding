package util

// https://www.codementor.io/tucnak/using-golang-for-competitive-programming-h8lhvXzt3

type Pair struct {
	X int
	Y int
}

type Pairs []Pair

func (r Pairs) Len() int           { return len(r) }
func (r Pairs) Less(i, j int) bool { return r[i].X < r[j].X || r[i].Y < r[j].Y }
func (r Pairs) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
