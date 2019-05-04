package dynamic

import (
	"testing"
)

func TestCoinchange(t *testing.T) {
	coins := []int{186, 419, 83, 408}
	amount := 6249

	fewest := coinChange(coins, amount)
	if fewest != 20 {
		t.FailNow()
	}
}
