package dynamic

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	n := 9
	expected := 34
	res := fibonacci(n)
	fmt.Printf("Fibonacci: %v %v\n", n, res)
	if res != expected {
		t.FailNow()
	}

	res = fibfunctor()(n)
	fmt.Printf("Fibonacci: %v %v\n", n, res)
	if res != expected {
		t.FailNow()
	}

	res = fibonacciN1(n)
	fmt.Printf("Fibonacci: %v %v\n", n, res)
	if res != expected {
		t.FailNow()
	}
}
