package recursion

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println(a)
	p := permute(a)
	fmt.Println("permute", p)
}

func TestPermute2(t *testing.T) {
	a := []int{1, 2, 3, 4}
	//fmt.Println(a)
	permute2(a, 0)
}
