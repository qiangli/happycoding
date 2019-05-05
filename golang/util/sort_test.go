package util

import (
	"sort"
	"testing"
)

func TestPairs(t *testing.T) {
	defer writer.Flush()

	items := Pairs{{4, 74.0}, {9, 34}, {9, 31}}
	sort.Sort(sort.Reverse(items))

	for _, item := range items {
		printf("(%d, %d)\n", item.X, item.Y)
	}
}
