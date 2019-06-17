package misc

import (
	"fmt"
	"testing"
)

func TestPermutationChar(t *testing.T) {
	permutationChar("ABC")
}

func TestCombination(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	combination(a, len(a), 3)
}

func TestCombinationNK(t *testing.T) {
	combinationNK(5, 3)
	fmt.Println("---")
	combinationNK2(5, 3)

}

func TestCompose(t *testing.T) {
	n := 5
	k := 3
	a := make([]int, n)
	compose(a, n, k, 0)
}

func TestCompose2(t *testing.T) {
	n := 5
	a := make([]int, n)
	compose2(a, n, 5, 0)
}
