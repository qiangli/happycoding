package string

import (
	"fmt"
	"testing"
)

func TestInteger2Words(t *testing.T) {
	n := 123
	fmt.Println(int2word(n))
	n = 12345
	fmt.Println(int2word(n))
	n = 1234567
	fmt.Println(int2word(n))
	fmt.Printf("%v\n", Max)
	n = 1234567891
	fmt.Println(int2word(n))
}
