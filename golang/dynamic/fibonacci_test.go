package dynamic

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	n := 9
	fmt.Printf("Fibonacci: %v %v\n", n, fibonacci(n))

	f := fibfunctor()
	fmt.Printf("Fibonacci: %v %v\n", n, f(n))
}
