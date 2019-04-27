package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := randomArray(15)
	fmt.Printf("input: %v\n", a)

	bubbleSort(a)

	fmt.Printf("result: %v\n", a)
}
