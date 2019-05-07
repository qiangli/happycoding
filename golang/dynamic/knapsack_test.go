package dynamic

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	value := []int{60, 100, 120}
	weight := []int{10, 20, 30}
	w := 50

	if knapsack(w, weight, value) != 220 {
		t.FailNow()
	}

	value = []int{10, 40, 30, 50}
	weight = []int{5, 4, 6, 3}
	w = 10
	knapsack(w, weight, value)
}
