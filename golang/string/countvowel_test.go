package string

import (
	"testing"
)

func TestCountVowel(t *testing.T) {
	s := []byte("geeksforgeeks portal")

	c := countVowel(s, len(s))
	if c != 7 {
		t.FailNow()
	}
}
