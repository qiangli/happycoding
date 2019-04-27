package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := randomArray(15)
	c := mergeSort(a)

	fmt.Printf("input: %v\n", a)
	fmt.Printf("result: %v\n", c)
}
