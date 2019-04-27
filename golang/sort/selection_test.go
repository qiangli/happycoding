package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := randomArray(15)
	fmt.Printf("input: %v\n", a)

	selectionSort(a)

	fmt.Printf("result: %v\n", a)
}
