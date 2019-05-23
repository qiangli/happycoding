package greedy

import (
	"fmt"
	"testing"
)

func TestFractionKnapsack(t *testing.T) {
	items := []Item{{60, 10}, {100, 20}, {120, 30}}
	fmt.Println(items)

	res := fractionKnapsack(50, items)
	fmt.Println("fraction knapsack", res)
}
