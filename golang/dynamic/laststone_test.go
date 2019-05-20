package dynamic

import (
	"fmt"
	"testing"
)

func TestLastStoneWeightII(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	v := lastStoneWeightII(stones)
	fmt.Println(v)

	// stones = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 14, 23, 37, 61, 98}
	// v = lastStoneWeightII(stones)
	// fmt.Println(v)

	// stones = []int{21, 16, 23, 32, 25, 13, 20, 18, 22, 21, 84, 35, 33, 17, 27, 24, 10, 19, 31, 26, 94, 37, 31, 25, 24, 25, 15, 23, 17, 13}
	// v = lastStoneWeightII(stones)
	// fmt.Println(v)
}
