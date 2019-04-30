package tree

import (
	"fmt"
	"testing"
)

func TestSameTree(t *testing.T) {
	t1 := &Tree{
		Value: 3,
		Left: &Tree{
			Value: 1,
			Left: &Tree{
				Value: 1,
			},
			Right: &Tree{
				Value: 2,
			},
		},
		Right: &Tree{
			Value: 8,
			Left: &Tree{
				Value: 5,
			},
			Right: &Tree{
				Value: 13,
			},
		},
	}
	t2 := &Tree{
		Value: 8,
		Left: &Tree{
			Value: 3,
			Left: &Tree{
				Value: 1,
				Left: &Tree{
					Value: 1,
				},
				Right: &Tree{
					Value: 2,
				},
			},
			Right: &Tree{
				Value: 5,
			},
		},
		Right: &Tree{
			Value: 13,
		},
	}

	fmt.Println(Same(t1, t2))
}
