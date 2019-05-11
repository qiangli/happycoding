package util

import (
	"fmt"
	"sort"
	"testing"
)

func TestPairs(t *testing.T) {
	items := Pairs{{4, 74.0}, {9, 34}, {9, 31}}
	sort.Sort(sort.Reverse(items))

	for _, item := range items {
		fmt.Printf("(%d, %d)\n", item.X, item.Y)
	}
}
