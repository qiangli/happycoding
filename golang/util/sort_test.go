package util

import (
	"sort"
	"testing"
)

func TestSortPair(t *testing.T) {
	defer writer.Flush()

	items := pairs{{4, 74.0}, {9, 34.56}, {9, 31.31}}
	sort.Sort(sort.Reverse(items))

	for _, item := range items {
		printf("(%d, %.2f)\n", item.x, item.y)
	}

	// (9, 34.56)
	// (9, 31.31)
	// (4, 74.00)
}
