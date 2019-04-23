package dynamic

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	factorial(9)

	f := facfunctor()
	t.Logf("factorial of 9: %v", f(9))
	t.Logf("factorial of 10: %v", f(10))
}
