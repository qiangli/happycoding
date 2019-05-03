package recursion

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	a := generateParenthesis(3)
	fmt.Println(a)
}
