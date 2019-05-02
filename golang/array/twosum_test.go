package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	a := []int{2, 7, 11, 15}
	target := 22

	s := twoSum(a, target)
	fmt.Println(s)
}
