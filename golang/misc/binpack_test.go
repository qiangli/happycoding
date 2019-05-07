package misc

import (
	"fmt"
	"testing"
)

func TestBinpack(t *testing.T) {
	weight := []int{2, 5, 4, 7, 1, 3, 8}
	c := 10
	fmt.Println(weight, c)
	binpack(weight, c)
}
